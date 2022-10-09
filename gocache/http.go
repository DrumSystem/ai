package gocache

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"gocache/consistenthash"
	pb "gocache/gocachepb"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

type HttpPool struct {
	self     string
	basePath string
	mu	sync.Mutex
	peers *consistenthash.Map
	httpGetter map[string]*httpGetter
}

type httpGetter struct {
	baseURL string
}

const (
	defaultBasePath = "/_gocache/"
	defaultReplicas = 50
)

func NewHttpPool(self string) *HttpPool {
	return &HttpPool{
		self:     self,
		basePath: defaultBasePath,
	}
}

func (p *HttpPool) Log(format string, v ...interface{}) {
	log.Printf("[Server %s] %s", p.self, fmt.Sprintf(format, v...))
}

func (p *HttpPool) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !strings.HasPrefix(r.URL.Path, p.basePath) {
		panic("HTTPPool serving unexpected path: " + r.URL.Path)
	}

	p.Log("%s %s", r.Method, r.URL.Path)
	// /<basepath>/<groupname>/<key> required
	parts := strings.SplitN(r.URL.Path[len(p.basePath):], "/", 2)
	if len(parts) != 2 {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	groupName := parts[0]
	key := parts[1]

	group := GetGroup(groupName)
	if group == nil {
		http.Error(w, "no such group: "+groupName, http.StatusNotFound)
		return
	}

	view, err := group.Get(key)
	body, err := proto.Marshal(&pb.Response{Value: view.ByteSlice()})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/octet-stream")
	//w.Write(view.ByteSlice())
	w.Write(body)
}

func (h *httpGetter) Get(in *pb.Request, out *pb.Response) error {
	u := fmt.Sprintf("%v%v/%v", h.baseURL, url.QueryEscape(in.GetGroup()), url.QueryEscape(in.GetKey()))
	log.Printf("URL PATh %v", u)
	resp, err := http.Get(u)
	if err != nil {
		return  err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Server returned: %v", resp.Status)
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	return nil, fmt.Errorf("reading response body:%v", err)
	//}
	if err = proto.Unmarshal(bytes, out); err != nil {
		return fmt.Errorf("decoding response body: %v", err)
	}

	return nil
}


// Set updates the pool's list of peers.
func (p *HttpPool) Set(peers ...string)  {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.peers = consistenthash.NewMap(defaultReplicas, nil)
	p.peers.Add(peers...)
	p.httpGetter = make(map[string]*httpGetter, len(peers))
	for _, peer := range peers {
		p.httpGetter[peer] = &httpGetter{
			baseURL: peer + p.basePath,
		}
	}
}

// PickPeer picks a peer according to key
func (p *HttpPool) PickPeer(key string) (PeerGetter, bool) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if peer := p.peers.Get(key); peer != "" && peer != p.self {
		p.Log("Pick peer %s", peer)
		return p.httpGetter[peer], true
	}
	return nil, false
}

var _ PeerGetter = (*httpGetter)(nil)

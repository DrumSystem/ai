package gocache

import pb "gocache/gocachepb"


// 根据对应的key找到对应的服务器ip

type PeerPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool)
}


// PeerGetter is the interface that must be implemented by a peer.
type PeerGetter interface {
	//Get(group string, key string) ([]byte, error)
	Get(in *pb.Request, out *pb.Response) error
}

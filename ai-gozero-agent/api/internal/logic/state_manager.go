package logic

import (
	"ai-gozero-agent/api/internal/svc"
	"ai-gozero-agent/api/internal/types"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"strings"
	"time"
)

const (
	statekeyPrefix = "chat_state:"
	stateTTl       = 24 * time.Hour
)

type StateManager struct {
	scvCtx *svc.ServiceContext
}

func NewStateManager(svcCtx *svc.ServiceContext) *StateManager {
	return &StateManager{
		scvCtx: svcCtx,
	}
}

func (sm *StateManager) GetOrInitState(chatId string) (string, error) {
	key := statekeyPrefix + chatId
	state, err := sm.scvCtx.Redis.Get(context.Background(), key).Result()
	if err == nil {
		return state, nil
	}

	if err == redis.Nil {
		if err := sm.scvCtx.Redis.Set(context.Background(), key, types.StateStart, stateTTl).Err(); err != nil {
			return types.StateStart, fmt.Errorf("初始化状态失败: %v", err)
		}
		return types.StateStart, nil
	}
	return types.StateStart, fmt.Errorf("获取状态失败: %v", err)
}

func (sm *StateManager) SetState(chatId, state string) error {
	key := statekeyPrefix + chatId
	if err := sm.scvCtx.Redis.Set(context.Background(), key, state, stateTTl).Err(); err != nil {
		return fmt.Errorf("设置状态失败: %v", err)
	}
	return nil

}

// 评估并更新状态（更智能的规则）
func (sm *StateManager) EvaluateAndUpdateState(chatId, aiResponse string) (string, error) {
	currentState, err := sm.GetOrInitState(chatId)
	if err != nil {
		return currentState, err
	}

	newState := sm.determineNewState(currentState, aiResponse)

	if newState != currentState {
		if err := sm.SetState(chatId, newState); err != nil {
			return currentState, err
		}
	}

	return newState, nil
}

// 状态转移决策逻辑
func (sm *StateManager) determineNewState(currentState, aiResponse string) string {
	lowerResponse := strings.ToLower(aiResponse)

	switch currentState {
	case types.StateStart:
		if containsAny(lowerResponse, []string{"你好", "欢迎", "面试开始"}) {
			return types.StateQuestion
		}

	case types.StateQuestion:
		if containsAny(lowerResponse, []string{"追问", "详细说明", "为什么", "怎么实现"}) {
			return types.StateFollowUp
		}
		if containsAny(lowerResponse, []string{"评估", "总结", "表现", "优缺点"}) {
			return types.StateEvaluate
		}

	case types.StateFollowUp:
		if containsAny(lowerResponse, []string{"评估", "总结", "表现", "优缺点"}) {
			return types.StateEvaluate
		}
		if containsAny(lowerResponse, []string{"下一个问题", "新问题"}) {
			return types.StateQuestion
		}

	case types.StateEvaluate:
		if containsAny(lowerResponse, []string{"结束", "再见", "感谢参加"}) {
			return types.StateEnd
		}
		if containsAny(lowerResponse, []string{"继续", "下一个问题"}) {
			return types.StateQuestion
		}

	case types.StateEnd:
		// 结束状态保持不变
	}

	return currentState
}

func containsAny(s string, substrs []string) bool {
	for _, sub := range substrs {
		if strings.Contains(s, sub) {
			return true
		}
	}
	return false
}

package server_code

import (
	"context"
	"plugin"
)

const (
	Success = "1000"
)

type Worker struct {
	/**
	 * 业务自定义参数结构
	 */
}

func (Worker) Do(ctx context.Context, plugin plugin.Plugin) (string, error) {
	/**
	 * 业务自定义处理逻辑
	 * 返回值为ErrorCode和错误信息
	 */
	return Success, nil
}

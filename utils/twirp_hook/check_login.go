package twirp_hook

import (
	"context"
	"github.com/learninto/sniper-api/utils/ctxkit"
	"github.com/learninto/sniper-api/utils/errors"

	"github.com/learninto/sniper-api/utils/twirp"
)

// NewCheckLogin 检查用户登录态，未登录直接报错返回
func NewCheckLogin() *twirp.ServerHooks {
	return &twirp.ServerHooks{
		RequestRouted: func(ctx context.Context) (context.Context, error) {
			if ctxkit.GetUserID(ctx) == 0 {
				return ctx, errors.NotLoginError
			}

			return ctx, nil
		},
	}
}

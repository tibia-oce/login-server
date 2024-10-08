package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tibia-oce/login-server/src/grpc/login_proto_messages"
)

func TestLoadSessionFromMessage(t *testing.T) {
	type args struct {
		sessionMsg *login_proto_messages.Session
	}
	tests := []struct {
		name string
		args args
		want Session
	}{{
		"is_not_premium",
		args{createSessionMessage(false)},
		Session{
			IsPremium:     false,
			PremiumUntil:  uint64(defaultNumber),
			SessionKey:    defaultString,
			LastLoginTime: defaultNumber,
			Status:        "active",
		}}, {
		"is_premium",
		args{createSessionMessage(true)},
		Session{
			IsPremium:     true,
			PremiumUntil:  uint64(defaultNumber),
			SessionKey:    defaultString,
			LastLoginTime: defaultNumber,
			Status:        "active",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, LoadSessionFromMessage(tt.args.sessionMsg))
		})
	}
}

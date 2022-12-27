// Copyright 2020 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package idtoken

import (
	"context"
	"reflect"
	"testing"

	"golang.org/x/oauth2"
	"google.golang.org/api/internal"
)

var TokenSource oauth2.TokenSource

func TestNewTokenSource(t *testing.T) {
	tests := []struct {
		name     string
		ctx      context.Context
		audience string
		want     oauth2.TokenSource
		wantErr  bool
	}{
		{
			name:     "works",
			ctx:      context.Background(),
			audience: "https://apikeys.googleapis.com",
			want:     TokenSource,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTokenSource(tt.ctx, tt.audience)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTokenSource() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			tok, err := got.Token()
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTokenSource() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			_, err = Validate(tt.ctx, tok.AccessToken, tt.audience)
			if err != nil {
				t.Errorf("NewTokenSource() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newTokenSource(t *testing.T) {
	type args struct {
		ctx      context.Context
		audience string
		ds       *internal.DialSettings
	}
	tests := []struct {
		name    string
		args    args
		want    oauth2.TokenSource
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newTokenSource(tt.args.ctx, tt.args.audience, tt.args.ds)
			if (err != nil) != tt.wantErr {
				t.Errorf("newTokenSource() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newTokenSource() = %v, want %v", got, tt.want)
			}
		})
	}
}

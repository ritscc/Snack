package env

import (
	"github.com/stretchr/testify/assert"
	"os"
	"strconv"
	"testing"
)

func setValue(key string, value string) error {
	if err := os.Setenv(key, value); err != nil {
		return err
	}
	return nil
}

func TestGetStrEnv(t *testing.T) {
	if err := setValue("SOME_KEY", "test"); err != nil {
		t.Fatal(err)
	}
	type args struct {
		key      string
		required bool
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// Test cases
		{
			name: "env variable exists",
			args: args{
				key:      "SOME_KEY",
				required: true,
			},
			want:    "test",
			wantErr: false,
		},
		{
			name: "env variable does not exist",
			args: args{
				key:      "NON_EXISTENT_KEY",
				required: true,
			},
			wantErr: true,
		},
		{
			name: "env variable does not exist and is not required",
			args: args{
				key:      "NON_EXISTENT_KEY",
				required: false,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetStrEnv(tt.args.key, tt.args.required)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestGetIntEnv(t *testing.T) {
	if err := setValue("SOME_KEY", strconv.Itoa(123)); err != nil {
		t.Fatal(err)
	}
	if err := setValue("INVALID_EXISTENT_KEY", "invalid"); err != nil {
		t.Fatal(err)
	}
	type args struct {
		key      string
		required bool
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "env variable exists",
			args: args{
				key:      "SOME_KEY",
				required: true,
			},
			want: 123,
		},
		{
			name: "env variable does not exist",
			args: args{
				key:      "NON_EXISTENT_KEY",
				required: true,
			},
			wantErr: true,
		},
		{
			name: "env variable does not exist and is not required",
			args: args{
				key:      "NON_EXISTENT_KEY",
				required: false,
			},
			wantErr: false,
		},
		{
			name: "env variable exist and is invalid",
			args: args{
				key:      "INVALID_EXISTENT_KEY",
				required: false,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetIntEnv(tt.args.key, tt.args.required)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

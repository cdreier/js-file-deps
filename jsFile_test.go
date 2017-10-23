package main

import "testing"

func Test_isOneLineImportStatement(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "test react import",
			args: args{line: "import React from 'react'"},
			want: true,
		},
		{
			name: "test simplest",
			args: args{line: "import from "},
			want: true,
		},
		{
			name: "test multiline",
			args: args{line: "import { "},
			want: false,
		},
		{
			name: "test all",
			args: args{line: "import {} from 'yaaay'"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isOneLineImportStatement(tt.args.line); got != tt.want {
				t.Errorf("isOneLineImportStatement() = %v, want %v", got, tt.want)
			}
		})
	}
}

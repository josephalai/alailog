package alailog

import (
	"os"
	"reflect"
	"testing"
)

func TestColor_Code(t *testing.T) {
	tests := []struct {
		name string
		c    Color
		want string
	}{
		{
			name: "Black",
			c:    Black,
			want: "30",
		},
		{
			name: "Red",
			c:    Red,
			want: "31",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Code(); got != tt.want {
				t.Errorf("Code() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestColor_String(t *testing.T) {
	tests := []struct {
		name string
		c    Color
		want string
	}{
		{name: "Black", c: Black, want: "\033[1;30m"}, {name: "Red", c: Red, want: "\033[1;31m"}, {name: "Green", c: Green, want: "\033[1;32m"}, {name: "Yellow", c: Yellow, want: "\033[1;33m"}, {name: "Blue", c: Blue, want: "\033[1;34m"}, {name: "Purple", c: Purple, want: "\033[1;35m"}, {name: "Cyan", c: Cyan, want: "\033[1;36m"}, {name: "White", c: White, want: "\033[1;37m"}, {name: "Reset", c: Reset, want: "\033[0m"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogger_Debug(t *testing.T) {
	type fields struct {
		file   *os.File
		level  Level
		stdout bool
		stderr bool
		color  bool
	}
	type args struct {
		message string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"", fields{nil, DebugLvl, false, false, false}, args{"test"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logger{
				file:   tt.fields.file,
				level:  tt.fields.level,
				stdout: tt.fields.stdout,
				stderr: tt.fields.stderr,
				color:  tt.fields.color,
			}
			l.Debug(tt.args.message)
		})
	}
}

func TestLogger_DebugBlack(t *testing.T) {
	type fields struct {
		file   *os.File
		level  Level
		stdout bool
		stderr bool
		color  bool
	}
	type args struct {
		message string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"", fields{nil, DebugLvl, false, false, false}, args{"test"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logger{
				file:   tt.fields.file,
				level:  tt.fields.level,
				stdout: tt.fields.stdout,
				stderr: tt.fields.stderr,
				color:  tt.fields.color,
			}
			l.DebugBlack(tt.args.message)
		})
	}
}

func TestLogger_DebugColor(t *testing.T) {
	type fields struct {
		file   *os.File
		level  Level
		stdout bool
		stderr bool
		color  bool
	}
	type args struct {
		color   Color
		message string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"", fields{nil, DebugLvl, true, false, true}, args{"test", "test"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logger{
				file:   tt.fields.file,
				level:  tt.fields.level,
				stdout: tt.fields.stdout,
				stderr: tt.fields.stderr,
				color:  tt.fields.color,
			}
			l.DebugColor(tt.args.color, tt.args.message)
		})
	}
}

func TestLogger_Error(t *testing.T) {
	type fields struct {
		file   *os.File
		level  Level
		stdout bool
		stderr bool
		color  bool
	}
	type args struct {
		message string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"", fields{nil, ErrorLvl, false, true, false}, args{"test"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logger{
				file:   tt.fields.file,
				level:  tt.fields.level,
				stdout: tt.fields.stdout,
				stderr: tt.fields.stderr,
				color:  tt.fields.color,
			}
			l.Error(tt.args.message)
		})
	}
}

func TestLogger_ErrorColor(t *testing.T) {
	type fields struct {
		file   *os.File
		level  Level
		stdout bool
		stderr bool
		color  bool
	}
	type args struct {
		color   Color
		message string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"", fields{nil, ErrorLvl, false, true, true}, args{"test", "test"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logger{
				file:   tt.fields.file,
				level:  tt.fields.level,
				stdout: tt.fields.stdout,
				stderr: tt.fields.stderr,
				color:  tt.fields.color,
			}
			l.ErrorColor(tt.args.color, tt.args.message)
		})
	}
}

func TestLogger_Fatal(t *testing.T) {
	type fields struct {
		file   *os.File
		level  Level
		stdout bool
		stderr bool
		color  bool
	}
	type args struct {
		message string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logger{
				file:   tt.fields.file,
				level:  tt.fields.level,
				stdout: tt.fields.stdout,
				stderr: tt.fields.stderr,
				color:  tt.fields.color,
			}
			l.Fatal(tt.args.message)
		})
	}
}

func TestLogger_FatalColor(t *testing.T) {
	type fields struct {
		file   *os.File
		level  Level
		stdout bool
		stderr bool
		color  bool
	}
	type args struct {
		color   Color
		message string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logger{
				file:   tt.fields.file,
				level:  tt.fields.level,
				stdout: tt.fields.stdout,
				stderr: tt.fields.stderr,
				color:  tt.fields.color,
			}
			l.FatalColor(tt.args.color, tt.args.message)
		})
	}
}

func TestLogger_Info(t *testing.T) {
	type fields struct {
		file   *os.File
		level  Level
		stdout bool
		stderr bool
		color  bool
	}
	type args struct {
		message string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logger{
				file:   tt.fields.file,
				level:  tt.fields.level,
				stdout: tt.fields.stdout,
				stderr: tt.fields.stderr,
				color:  tt.fields.color,
			}
			l.Info(tt.args.message)
		})
	}
}

func TestLogger_InfoColor(t *testing.T) {
	type fields struct {
		file   *os.File
		level  Level
		stdout bool
		stderr bool
		color  bool
	}
	type args struct {
		color   Color
		message string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logger{
				file:   tt.fields.file,
				level:  tt.fields.level,
				stdout: tt.fields.stdout,
				stderr: tt.fields.stderr,
				color:  tt.fields.color,
			}
			l.InfoColor(tt.args.color, tt.args.message)
		})
	}
}

func TestLogger_Log(t *testing.T) {
	type fields struct {
		file   *os.File
		level  Level
		stdout bool
		stderr bool
		color  bool
	}
	type args struct {
		level   Level
		message string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logger{
				file:   tt.fields.file,
				level:  tt.fields.level,
				stdout: tt.fields.stdout,
				stderr: tt.fields.stderr,
				color:  tt.fields.color,
			}
			l.Log(tt.args.level, tt.args.message)
		})
	}
}

func TestLogger_LogColor(t *testing.T) {
	type fields struct {
		file   *os.File
		level  Level
		stdout bool
		stderr bool
		color  bool
	}
	type args struct {
		level   Level
		color   Color
		message string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logger{
				file:   tt.fields.file,
				level:  tt.fields.level,
				stdout: tt.fields.stdout,
				stderr: tt.fields.stderr,
				color:  tt.fields.color,
			}
			l.LogColor(tt.args.level, tt.args.color, tt.args.message)
		})
	}
}

func TestLogger_SetLevel(t *testing.T) {
	type fields struct {
		file   *os.File
		level  Level
		stdout bool
		stderr bool
		color  bool
	}
	type args struct {
		level Level
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logger{
				file:   tt.fields.file,
				level:  tt.fields.level,
				stdout: tt.fields.stdout,
				stderr: tt.fields.stderr,
				color:  tt.fields.color,
			}
			l.SetLevel(tt.args.level)
		})
	}
}

func TestLogger_SetStderr(t *testing.T) {
	type fields struct {
		file   *os.File
		level  Level
		stdout bool
		stderr bool
		color  bool
	}
	type args struct {
		stderr bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logger{
				file:   tt.fields.file,
				level:  tt.fields.level,
				stdout: tt.fields.stdout,
				stderr: tt.fields.stderr,
				color:  tt.fields.color,
			}
			l.SetStderr(tt.args.stderr)
		})
	}
}

func TestLogger_SetStdout(t *testing.T) {
	type fields struct {
		file   *os.File
		level  Level
		stdout bool
		stderr bool
		color  bool
	}
	type args struct {
		stdout bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logger{
				file:   tt.fields.file,
				level:  tt.fields.level,
				stdout: tt.fields.stdout,
				stderr: tt.fields.stderr,
				color:  tt.fields.color,
			}
			l.SetStdout(tt.args.stdout)
		})
	}
}

func TestLogger_Warn(t *testing.T) {
	type fields struct {
		file   *os.File
		level  Level
		stdout bool
		stderr bool
		color  bool
	}
	type args struct {
		message string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logger{
				file:   tt.fields.file,
				level:  tt.fields.level,
				stdout: tt.fields.stdout,
				stderr: tt.fields.stderr,
				color:  tt.fields.color,
			}
			l.Warn(tt.args.message)
		})
	}
}

func TestLogger_WarnColor(t *testing.T) {
	type fields struct {
		file   *os.File
		level  Level
		stdout bool
		stderr bool
		color  bool
	}
	type args struct {
		color   Color
		message string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"", fields{nil, WarnLvl, true, false, true}, args{"test", "test"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logger{
				file:   tt.fields.file,
				level:  tt.fields.level,
				stdout: tt.fields.stdout,
				stderr: tt.fields.stderr,
				color:  tt.fields.color,
			}
			l.WarnColor(tt.args.color, tt.args.message)
		})
	}
}

func TestNewLogger(t *testing.T) {
	type args struct {
		file   *os.File
		level  Level
		stdout bool
		stderr bool
		color  bool
	}
	tests := []struct {
		name string
		args args
		want *Logger
	}{
		{
			name: "Test",
			args: args{
				file:   nil,
				level:  DebugLvl,
				stdout: true,
				stderr: true,
				color:  true,
			},
			want: NewLogger(nil, DebugLvl, true, true, true),
		},
		{
			name: "Test",
			args: args{
				file:   nil,
				level:  DebugLvl,
				stdout: true,
				stderr: true,
				color:  true,
			},
			want: &Logger{
				file:   nil,
				level:  DebugLvl,
				stdout: true,
				stderr: true,
				color:  true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLogger(tt.args.file, tt.args.level, tt.args.stdout, tt.args.stderr, tt.args.color); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLogger() = %v, want %v", got, tt.want)
			}
		})
	}
}

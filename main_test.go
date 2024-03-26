package main

import "testing"

func Test_checkSuppose(t *testing.T) {
	type args struct {
		goal    []color
		suppose []color
	}
	var goal = []color{
		pink, blue, orange, green,
	}

	tests := []struct {
		name   string
		args   args
		blacks int8
		whites int8
		isWin  bool
	}{
		{
			name: "Only 2 white",
			args: args{
				goal:    goal,
				suppose: []color{lightGreen, pink, brown, orange},
			},
			blacks: 0,
			whites: 2,
			isWin:  false,
		}, {
			name: "Only 2 blacks",
			args: args{
				goal:    goal,
				suppose: []color{pink, blue, lightGreen, brown},
			},
			blacks: 2,
			whites: 0,
			isWin:  false,
		}, {
			name: "4 blacks",
			args: args{
				goal:    goal,
				suppose: []color{pink, blue, orange, green},
			},
			blacks: 4,
			whites: 0,
			isWin:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := checkSuppose(tt.args.goal, tt.args.suppose)
			if got != tt.blacks {
				t.Errorf("checkSuppose() suppose = %v, blacks %v", got, tt.blacks)
			}
			if got1 != tt.whites {
				t.Errorf("checkSuppose() got1 = %v, blacks %v", got1, tt.whites)
			}
			if got2 != tt.isWin {
				t.Errorf("checkSuppose() got2 = %v, blacks %v", got2, tt.isWin)
			}
		})
	}
}

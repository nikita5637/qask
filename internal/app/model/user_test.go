package model

import "testing"

func TestUser_Validate(t *testing.T) {
	tests := []struct {
		name    string
		user    User
		wantErr bool
	}{
		{
			name: "valid user",
			user: User{
				userPublic{
					FirstName: "validFirstName",
				},
				userPrivate{
					UserName: "validUserName",
					TgID:     2,
				},
			},
			wantErr: false,
		},
		{
			name: "invalid user(invalid FirstName)",
			user: User{
				userPublic{
					FirstName: "no",
				},
				userPrivate{
					UserName: "validUserName",
					TgID:     2,
				},
			},
			wantErr: true,
		},
		{
			name: "invalid user(invalid UserName)",
			user: User{
				userPublic{
					FirstName: "validUserName",
				},
				userPrivate{
					UserName: "no",
					TgID:     2,
				},
			},
			wantErr: true,
		},
		{
			name: "invalid user(invalid Telegram ID(eq 0))",
			user: User{
				userPublic{
					FirstName: "validFirstName",
				},
				userPrivate{
					UserName: "validUserName",
					TgID:     0,
				},
			},
			wantErr: true,
		},
		{
			name: "invalid user(invalid Telegram ID(less than 0))",
			user: User{
				userPublic{
					FirstName: "validFirstName",
				},
				userPrivate{
					UserName: "validUserName",
					TgID:     -1,
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.user.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("User.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

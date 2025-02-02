package usecase

import (
	"context"
	"database/sql"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/h-hiwatashi/go-practice-sample/db"
	"github.com/h-hiwatashi/go-practice-sample/models"

	_ "github.com/go-sql-driver/mysql"
)

func TestAdd(t *testing.T) {
	// テスト用のDB接続取得
	testDB := db.NewUnitTestDB()

	// Inputの情報で作成した後に、作成されたユーザーIDを元にユーザーを取得して、Wantと比較する。
	cases := []struct {
		label string
		Input AddUserInput
		Want  *models.User
	}{
		{
			"ユーザー登録のテスト",
			AddUserInput{
				Name: "John",
				Age:  40,
			},
			&models.User{
				Name: "John",
				Age:  40,
			},
		},
	}

	for _, c := range cases {
		t.Run(c.label, func(t *testing.T) {
			testDB.InTx(func(ctx context.Context, tx *sql.Tx) {
				usecase := NewUserUsecase(ctx, tx)
				output, err := usecase.Add(c.Input)
				if err != nil {
					t.Fatal("error:", err)
				}

				user, _ := models.FindUser(ctx, tx, output.ID)

				// IDは何が入るかわからないので、比較する絡むから外す。
				opts := []cmp.Option{
					cmpopts.IgnoreFields(models.User{}, "ID"),
				}

				// 作成したユーザーと、求める結果の比較をしている。
				// 差があった場合にエラーにしている。
				if diff := cmp.Diff(user, c.Want, opts...); diff != "" {
					t.Errorf(diff)
				}
			})
		})
	}
}

func TestGet(t *testing.T) {
	// テスト用のDB接続取得
	testDB := db.NewUnitTestDB()

	// Getした結果とWantの情報を比較する
	cases := []struct {
		label string
		Input GetUserInput
		Want  GetUserOutput
	}{
		{
			"ユーザー情報取得のテスト",
			GetUserInput{
				ID: 4,
			},
			GetUserOutput{
				ID:   4,
				Name: "taro",
				Age:  25,
			},
		},
	}

	for _, c := range cases {
		t.Run(c.label, func(t *testing.T) {
			testDB.InTx(func(ctx context.Context, tx *sql.Tx) {
				usecase := NewUserUsecase(ctx, tx)
				output, err := usecase.Get(c.Input)
				if err != nil {
					t.Fatal("error:", err)
				}

				// 取得したユーザーと、求める結果の比較をしている。
				// 差があった場合にエラーにしている。
				if diff := cmp.Diff(output, c.Want); diff != "" {
					t.Errorf(diff)
				}
			})
		})
	}
}

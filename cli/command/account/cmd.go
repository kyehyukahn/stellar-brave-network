package account

import (
	"github.com/soonkuk/stellar-brave-network/cli/command"
	"github.com/spf13/cobra"
	"fmt"
	"github.com/stellar/go/clients/horizon"
)

func NewAccountCommand(cli *command.BraveCli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "account",
		Short: "account",
		Long:  "",
	}

	balance := &cobra.Command{
		Use:   "balance",
		Short: "balance",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			alias := cmd.Flags().Lookup("alias").Value.String()
			address := cmd.Flags().Lookup("address").Value.String()

			if alias != "" {
				address = cli.Account[alias]
			} else if address == "" {
				cmd.Usage()
				return
			}

			account, err := cli.HorizonClient1().LoadAccount(address)

			if err == nil {
				fmt.Println("My account address:", account.AccountID)
				for _, v := range account.Balances {
					fmt.Printf("type: %s balance: %s\n", v.Asset.Type, v.Balance)
				}
			} else {
				horizonErr := err.(*horizon.Error)
				if horizonErr.Response.StatusCode == 404 {
					fmt.Println("user not found")
				}
				panic(err)
			}
		},
	}

	balance.Flags().String("alias", "", "config.yaml 설정된 계정 alias")
	balance.Flags().String("address", "", "계좌 주소")

	balance2 := &cobra.Command{
		Use:   "balance2",
		Short: "balance with 2 horizons",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			alias := cmd.Flags().Lookup("alias").Value.String()
			address := cmd.Flags().Lookup("address").Value.String()

			if alias != "" {
				address = cli.Account[alias]
			} else if address == "" {
				cmd.Usage()
				return
			}

			account, err := cli.HorizonClient1().LoadAccount(address)

			if err == nil {
				fmt.Println("My account address:", account.AccountID)
				for _, v := range account.Balances {
					fmt.Printf("type: %s balance: %s\n", v.Asset.Type, v.Balance)
				}
			} else {
				horizonErr := err.(*horizon.Error)
				if horizonErr.Response.StatusCode == 404 {
					fmt.Println("user not found")
				}
				panic(err)
			}
			account2, err2 := cli.HorizonClient2().LoadAccount(address)

			if err == nil {
				fmt.Println("My account address:", account2.AccountID)
				for _, v := range account2.Balances {
					fmt.Printf("type: %s balance: %s\n", v.Asset.Type, v.Balance)
				}
			} else {
				horizonErr := err.(*horizon.Error)
				if horizonErr.Response.StatusCode == 404 {
					fmt.Println("user not found")
				}
				panic(err2)
			}

		},
	}

	balance2.Flags().String("alias", "", "config.yaml 설정된 계정 alias")
	balance2.Flags().String("address", "", "계좌 주소")

	balance3 := &cobra.Command{
		Use:   "balance3",
		Short: "balance with 3 horizons",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			alias := cmd.Flags().Lookup("alias").Value.String()
			address := cmd.Flags().Lookup("address").Value.String()

			if alias != "" {
				address = cli.Account[alias]
			} else if address == "" {
				cmd.Usage()
				return
			}

			account, err := cli.HorizonClient1().LoadAccount(address)

			if err == nil {
				fmt.Println("My account address:", account.AccountID)
				for _, v := range account.Balances {
					fmt.Printf("type: %s balance: %s\n", v.Asset.Type, v.Balance)
				}
			} else {
				horizonErr := err.(*horizon.Error)
				if horizonErr.Response.StatusCode == 404 {
					fmt.Println("user not found")
				}
				panic(err)
			}
			account2, err2 := cli.HorizonClient2().LoadAccount(address)

			if err == nil {
				fmt.Println("My account address:", account2.AccountID)
				for _, v := range account2.Balances {
					fmt.Printf("type: %s balance: %s\n", v.Asset.Type, v.Balance)
				}
			} else {
				horizonErr := err2.(*horizon.Error)
				if horizonErr.Response.StatusCode == 404 {
					fmt.Println("user not found")
				}
				panic(err2)
			}

			account3, err3 := cli.HorizonClient3().LoadAccount(address)

			if err == nil {
				fmt.Println("My account address:", account3.AccountID)
				for _, v := range account3.Balances {
					fmt.Printf("type: %s balance: %s\n", v.Asset.Type, v.Balance)
				}
			} else {
				horizonErr := err3.(*horizon.Error)
				if horizonErr.Response.StatusCode == 404 {
					fmt.Println("user not found")
				}
				panic(err3)
			}
		},
	}

	balance3.Flags().String("alias", "", "config.yaml 설정된 계정 alias")
	balance3.Flags().String("address", "", "계좌 주소")
	cmd.AddCommand(balance, balance2, balance3)

	return cmd
}

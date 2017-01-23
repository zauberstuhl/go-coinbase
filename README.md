# Coinbase Golang API Library

This one is tested against APIv2.

## Example

    c := coinbase.APIClient{
      Key: os.Getenv("COINBASE_KEY"),
      Secret: os.Getenv("COINBASE_SECRET"),
    }
    accounts, err := c.Accounts()
    if err != nil {
      fmt.Println(err)
      return
    }

    for i := range accounts.Data {
      acc := accounts.Data[i]
      fmt.Printf("%s %s %s %f %s\n",
        acc.Id, acc.Name, acc.Type,
        acc.Balance.Amount, acc.Balance.Currency)
    }


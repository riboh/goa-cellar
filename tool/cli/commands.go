package cli

import (
	"encoding/json"
	"fmt"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa-cellar/client"
	goaclient "github.com/goadesign/goa/client"
	uuid "github.com/goadesign/goa/uuid"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

type (
	// CreateAccountCommand is the command line data structure for the create action of account
	CreateAccountCommand struct {
		Payload     string
		ContentType string
		PrettyPrint bool
	}

	// DeleteAccountCommand is the command line data structure for the delete action of account
	DeleteAccountCommand struct {
		// Account ID
		AccountID   int
		PrettyPrint bool
	}

	// ListAccountCommand is the command line data structure for the list action of account
	ListAccountCommand struct {
		PrettyPrint bool
	}

	// ShowAccountCommand is the command line data structure for the show action of account
	ShowAccountCommand struct {
		// Account ID
		AccountID   int
		PrettyPrint bool
	}

	// UpdateAccountCommand is the command line data structure for the update action of account
	UpdateAccountCommand struct {
		Payload     string
		ContentType string
		// Account ID
		AccountID   int
		PrettyPrint bool
	}

	// CreateBottleCommand is the command line data structure for the create action of bottle
	CreateBottleCommand struct {
		Payload     string
		ContentType string
		// Account ID
		AccountID   int
		PrettyPrint bool
	}

	// DeleteBottleCommand is the command line data structure for the delete action of bottle
	DeleteBottleCommand struct {
		// Account ID
		AccountID   int
		BottleID    int
		PrettyPrint bool
	}

	// ListBottleCommand is the command line data structure for the list action of bottle
	ListBottleCommand struct {
		// Account ID
		AccountID int
		// Filter by years
		Years       []int
		PrettyPrint bool
	}

	// RateBottleCommand is the command line data structure for the rate action of bottle
	RateBottleCommand struct {
		Payload     string
		ContentType string
		// Account ID
		AccountID   int
		BottleID    int
		PrettyPrint bool
	}

	// ShowBottleCommand is the command line data structure for the show action of bottle
	ShowBottleCommand struct {
		// Account ID
		AccountID   int
		BottleID    int
		PrettyPrint bool
	}

	// UpdateBottleCommand is the command line data structure for the update action of bottle
	UpdateBottleCommand struct {
		Payload     string
		ContentType string
		// Account ID
		AccountID   int
		BottleID    int
		PrettyPrint bool
	}

	// WatchBottleCommand is the command line data structure for the watch action of bottle
	WatchBottleCommand struct {
		// Account ID
		AccountID   int
		BottleID    int
		PrettyPrint bool
	}

	// HealthHealthCommand is the command line data structure for the health action of health
	HealthHealthCommand struct {
		PrettyPrint bool
	}

	// DownloadCommand is the command line data structure for the download command.
	DownloadCommand struct {
		// OutFile is the path to the download output file.
		OutFile string
	}
)

// RegisterCommands registers the resource action CLI commands.
func RegisterCommands(app *cobra.Command, c *client.Client) {
	var command, sub *cobra.Command
	command = &cobra.Command{
		Use:   "create",
		Short: `create action`,
	}
	tmp1 := new(CreateAccountCommand)
	sub = &cobra.Command{
		Use:   `account ["/cellar/accounts"]`,
		Short: ``,
		Long: `

Payload example:

{
   "name": "test"
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp1.Run(c, args) },
	}
	tmp1.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp1.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp2 := new(CreateBottleCommand)
	sub = &cobra.Command{
		Use:   `bottle ["/cellar/accounts/ACCOUNTID/bottles"]`,
		Short: ``,
		Long: `

Payload example:

{
   "color": "white",
   "country": "USA",
   "name": "Number 8",
   "region": "Napa Valley",
   "review": "Great and inexpensive",
   "sweetness": 1,
   "varietal": "Merlot",
   "vineyard": "Asti",
   "vintage": 2012
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp2.Run(c, args) },
	}
	tmp2.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp2.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "delete",
		Short: `delete action`,
	}
	tmp3 := new(DeleteAccountCommand)
	sub = &cobra.Command{
		Use:   `account ["/cellar/accounts/ACCOUNTID"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp3.Run(c, args) },
	}
	tmp3.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp3.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp4 := new(DeleteBottleCommand)
	sub = &cobra.Command{
		Use:   `bottle ["/cellar/accounts/ACCOUNTID/bottles/BOTTLEID"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp4.Run(c, args) },
	}
	tmp4.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp4.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "health",
		Short: `Perform health check.`,
	}
	tmp5 := new(HealthHealthCommand)
	sub = &cobra.Command{
		Use:   `health ["/cellar/_ah/health"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp5.Run(c, args) },
	}
	tmp5.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp5.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "list",
		Short: `list action`,
	}
	tmp6 := new(ListAccountCommand)
	sub = &cobra.Command{
		Use:   `account ["/cellar/accounts"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp6.Run(c, args) },
	}
	tmp6.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp6.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp7 := new(ListBottleCommand)
	sub = &cobra.Command{
		Use:   `bottle ["/cellar/accounts/ACCOUNTID/bottles"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp7.Run(c, args) },
	}
	tmp7.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp7.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "rate",
		Short: ``,
	}
	tmp8 := new(RateBottleCommand)
	sub = &cobra.Command{
		Use:   `bottle ["/cellar/accounts/ACCOUNTID/bottles/BOTTLEID/actions/rate"]`,
		Short: ``,
		Long: `

Payload example:

{
   "rating": 3
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp8.Run(c, args) },
	}
	tmp8.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp8.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "show",
		Short: `show action`,
	}
	tmp9 := new(ShowAccountCommand)
	sub = &cobra.Command{
		Use:   `account ["/cellar/accounts/ACCOUNTID"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp9.Run(c, args) },
	}
	tmp9.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp9.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp10 := new(ShowBottleCommand)
	sub = &cobra.Command{
		Use:   `bottle ["/cellar/accounts/ACCOUNTID/bottles/BOTTLEID"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp10.Run(c, args) },
	}
	tmp10.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp10.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "update",
		Short: `update action`,
	}
	tmp11 := new(UpdateAccountCommand)
	sub = &cobra.Command{
		Use:   `account ["/cellar/accounts/ACCOUNTID"]`,
		Short: ``,
		Long: `

Payload example:

{
   "name": "test"
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp11.Run(c, args) },
	}
	tmp11.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp11.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp12 := new(UpdateBottleCommand)
	sub = &cobra.Command{
		Use:   `bottle ["/cellar/accounts/ACCOUNTID/bottles/BOTTLEID"]`,
		Short: ``,
		Long: `

Payload example:

{
   "color": "white",
   "country": "USA",
   "name": "Number 8",
   "region": "Napa Valley",
   "review": "Great and inexpensive",
   "sweetness": 1,
   "varietal": "Merlot",
   "vineyard": "Asti",
   "vintage": 2012
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp12.Run(c, args) },
	}
	tmp12.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp12.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "watch",
		Short: `Retrieve bottle with given id`,
	}
	tmp13 := new(WatchBottleCommand)
	sub = &cobra.Command{
		Use:   `bottle ["/cellar/accounts/ACCOUNTID/bottles/BOTTLEID/watch"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp13.Run(c, args) },
	}
	tmp13.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp13.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)

	dl := new(DownloadCommand)
	dlc := &cobra.Command{
		Use:   "download [PATH]",
		Short: "Download file with given path",
		RunE: func(cmd *cobra.Command, args []string) error {
			return dl.Run(c, args)
		},
	}
	dlc.Flags().StringVar(&dl.OutFile, "out", "", "Output file")
	app.AddCommand(dlc)
}

func intFlagVal(name string, parsed int) *int {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func float64FlagVal(name string, parsed float64) *float64 {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func boolFlagVal(name string, parsed bool) *bool {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func stringFlagVal(name string, parsed string) *string {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func hasFlag(name string) bool {
	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "--"+name) {
			return true
		}
	}
	return false
}

func jsonVal(val string) (*interface{}, error) {
	var t interface{}
	err := json.Unmarshal([]byte(val), &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func jsonArray(ins []string) ([]interface{}, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []interface{}
	for _, id := range ins {
		val, err := jsonVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, val)
	}
	return vals, nil
}

func timeVal(val string) (*time.Time, error) {
	t, err := time.Parse("RFC3339", val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func timeArray(ins []string) ([]time.Time, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []time.Time
	for _, id := range ins {
		val, err := timeVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func uuidVal(val string) (*uuid.UUID, error) {
	t, err := uuid.FromString(val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func uuidArray(ins []string) ([]uuid.UUID, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []uuid.UUID
	for _, id := range ins {
		val, err := uuidVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func float64Val(val string) (*float64, error) {
	t, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func float64Array(ins []string) ([]float64, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []float64
	for _, id := range ins {
		val, err := float64Val(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func boolVal(val string) (*bool, error) {
	t, err := strconv.ParseBool(val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func boolArray(ins []string) ([]bool, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []bool
	for _, id := range ins {
		val, err := boolVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

// Run makes the HTTP request corresponding to the CreateAccountCommand command.
func (cmd *CreateAccountCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/cellar/accounts"
	}
	var payload client.CreateAccountPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.CreateAccount(ctx, path, &payload, cmd.ContentType)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *CreateAccountCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
}

// Run makes the HTTP request corresponding to the DeleteAccountCommand command.
func (cmd *DeleteAccountCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/cellar/accounts/%v", cmd.AccountID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.DeleteAccount(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *DeleteAccountCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var accountID int
	cc.Flags().IntVar(&cmd.AccountID, "accountID", accountID, `Account ID`)
}

// Run makes the HTTP request corresponding to the ListAccountCommand command.
func (cmd *ListAccountCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/cellar/accounts"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ListAccount(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ListAccountCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
}

// Run makes the HTTP request corresponding to the ShowAccountCommand command.
func (cmd *ShowAccountCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/cellar/accounts/%v", cmd.AccountID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ShowAccount(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ShowAccountCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var accountID int
	cc.Flags().IntVar(&cmd.AccountID, "accountID", accountID, `Account ID`)
}

// Run makes the HTTP request corresponding to the UpdateAccountCommand command.
func (cmd *UpdateAccountCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/cellar/accounts/%v", cmd.AccountID)
	}
	var payload client.UpdateAccountPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.UpdateAccount(ctx, path, &payload, cmd.ContentType)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *UpdateAccountCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
	var accountID int
	cc.Flags().IntVar(&cmd.AccountID, "accountID", accountID, `Account ID`)
}

// Run makes the HTTP request corresponding to the CreateBottleCommand command.
func (cmd *CreateBottleCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/cellar/accounts/%v/bottles", cmd.AccountID)
	}
	var payload client.CreateBottlePayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.CreateBottle(ctx, path, &payload, cmd.ContentType)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *CreateBottleCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
	var accountID int
	cc.Flags().IntVar(&cmd.AccountID, "accountID", accountID, `Account ID`)
}

// Run makes the HTTP request corresponding to the DeleteBottleCommand command.
func (cmd *DeleteBottleCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/cellar/accounts/%v/bottles/%v", cmd.AccountID, cmd.BottleID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.DeleteBottle(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *DeleteBottleCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var accountID int
	cc.Flags().IntVar(&cmd.AccountID, "accountID", accountID, `Account ID`)
	var bottleID int
	cc.Flags().IntVar(&cmd.BottleID, "bottleID", bottleID, ``)
}

// Run makes the HTTP request corresponding to the ListBottleCommand command.
func (cmd *ListBottleCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/cellar/accounts/%v/bottles", cmd.AccountID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ListBottle(ctx, path, cmd.Years)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ListBottleCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var accountID int
	cc.Flags().IntVar(&cmd.AccountID, "accountID", accountID, `Account ID`)
	var years []int
	cc.Flags().IntSliceVar(&cmd.Years, "years", years, `Filter by years`)
}

// Run makes the HTTP request corresponding to the RateBottleCommand command.
func (cmd *RateBottleCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/cellar/accounts/%v/bottles/%v/actions/rate", cmd.AccountID, cmd.BottleID)
	}
	var payload client.RateBottlePayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.RateBottle(ctx, path, &payload, cmd.ContentType)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *RateBottleCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
	var accountID int
	cc.Flags().IntVar(&cmd.AccountID, "accountID", accountID, `Account ID`)
	var bottleID int
	cc.Flags().IntVar(&cmd.BottleID, "bottleID", bottleID, ``)
}

// Run makes the HTTP request corresponding to the ShowBottleCommand command.
func (cmd *ShowBottleCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/cellar/accounts/%v/bottles/%v", cmd.AccountID, cmd.BottleID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ShowBottle(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ShowBottleCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var accountID int
	cc.Flags().IntVar(&cmd.AccountID, "accountID", accountID, `Account ID`)
	var bottleID int
	cc.Flags().IntVar(&cmd.BottleID, "bottleID", bottleID, ``)
}

// Run makes the HTTP request corresponding to the UpdateBottleCommand command.
func (cmd *UpdateBottleCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/cellar/accounts/%v/bottles/%v", cmd.AccountID, cmd.BottleID)
	}
	var payload client.BottlePayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.UpdateBottle(ctx, path, &payload, cmd.ContentType)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *UpdateBottleCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
	var accountID int
	cc.Flags().IntVar(&cmd.AccountID, "accountID", accountID, `Account ID`)
	var bottleID int
	cc.Flags().IntVar(&cmd.BottleID, "bottleID", bottleID, ``)
}

// Run establishes a websocket connection for the WatchBottleCommand command.
func (cmd *WatchBottleCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/cellar/accounts/%v/bottles/%v/watch", cmd.AccountID, cmd.BottleID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	ws, err := c.WatchBottle(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}
	go goaclient.WSWrite(ws)
	goaclient.WSRead(ws)

	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *WatchBottleCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var accountID int
	cc.Flags().IntVar(&cmd.AccountID, "accountID", accountID, `Account ID`)
	var bottleID int
	cc.Flags().IntVar(&cmd.BottleID, "bottleID", bottleID, ``)
}

// Run makes the HTTP request corresponding to the HealthHealthCommand command.
func (cmd *HealthHealthCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/cellar/_ah/health"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.HealthHealth(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *HealthHealthCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
}

// Run downloads files with given paths.
func (cmd *DownloadCommand) Run(c *client.Client, args []string) error {
	var (
		fnf func(context.Context, string) (int64, error)
		fnd func(context.Context, string, string) (int64, error)

		rpath   = args[0]
		outfile = cmd.OutFile
		logger  = goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
		ctx     = goa.WithLogger(context.Background(), logger)
		err     error
	)

	if rpath[0] != '/' {
		rpath = "/" + rpath
	}
	if strings.HasPrefix(rpath, "/js/") {
		fnd = c.DownloadJs
		rpath = rpath[4:]
		if outfile == "" {
			_, outfile = path.Split(rpath)
		}
		goto found
	}
	return fmt.Errorf("don't know how to download %s", rpath)
found:
	ctx = goa.WithLogContext(ctx, "file", outfile)
	if fnf != nil {
		_, err = fnf(ctx, outfile)
	} else {
		_, err = fnd(ctx, rpath, outfile)
	}
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	return nil
}

// Run downloads files with given paths.
func (cmd *DownloadCommand) Run(c *client.Client, args []string) error {
	var (
		fnf func(context.Context, string) (int64, error)
		fnd func(context.Context, string, string) (int64, error)

		rpath   = args[0]
		outfile = cmd.OutFile
		logger  = goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
		ctx     = goa.WithLogger(context.Background(), logger)
		err     error
	)

	if rpath[0] != '/' {
		rpath = "/" + rpath
	}
	if rpath == "/ui" {
		fnf = c.DownloadUI
		if outfile == "" {
			outfile = "index.html"
		}
		goto found
	}
	return fmt.Errorf("don't know how to download %s", rpath)
found:
	ctx = goa.WithLogContext(ctx, "file", outfile)
	if fnf != nil {
		_, err = fnf(ctx, outfile)
	} else {
		_, err = fnd(ctx, rpath, outfile)
	}
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	return nil
}

// Run downloads files with given paths.
func (cmd *DownloadCommand) Run(c *client.Client, args []string) error {
	var (
		fnf func(context.Context, string) (int64, error)
		fnd func(context.Context, string, string) (int64, error)

		rpath   = args[0]
		outfile = cmd.OutFile
		logger  = goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
		ctx     = goa.WithLogger(context.Background(), logger)
		err     error
	)

	if rpath[0] != '/' {
		rpath = "/" + rpath
	}
	if rpath == "/swagger.json" {
		fnf = c.DownloadSwaggerJSON
		if outfile == "" {
			outfile = "swagger.json"
		}
		goto found
	}
	return fmt.Errorf("don't know how to download %s", rpath)
found:
	ctx = goa.WithLogContext(ctx, "file", outfile)
	if fnf != nil {
		_, err = fnf(ctx, outfile)
	} else {
		_, err = fnd(ctx, rpath, outfile)
	}
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	return nil
}

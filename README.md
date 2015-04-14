<h1>YQuotes</h1>
<p>Simple way to get stock quotes from Yahoo Finance.</p>
<p>
  <ul>
    <li>Get stock information (price, name and etc.)</li>
    <li>Get historical price data</li>
  </ul>
</p>

<h2>Install</h2>
<code>go get github.com/doneland/yquotes</code>

<h2>How to use</h2>
<h3>Get price information of single stock</h3>
<p>
  Client can get price information about stock from Finance Yahoo by calling 
  <code>NewStock(symbol string, history bool)</code> method. It retruns
  <code>Stock</code> type with recent price information. <code>history</code> property 
  directs wether historical data should be loaded or not.
</p>

<pre>
  // Get stock information without historical data. If you want to load historical
  // data, second argument to <code>TRUE</code>.
  stock, err := yquotes.NewStock("AAPL", false)
  if err != nil {
    // handle error
  }

  symbol := stock.Symbol // AAPL
  name := stock.Name // Apple Inc.
  
  // Price information
  price     := stock.Price // Price struct 
  bid       := price.Bid
  ask       := price.Ask
  open      := price.Open
  prevClose := price.PreviousClose
  last      := price.Last
  date      := price.Date 
</pre>

<h3>Get historical information</h3>
<h4>History for selected number of years</h4>

<p>
  Function <code>HistoryForYears</code> accepts three parameters: symbol, number of 
  years and frequency (daily, monthly). Frequency is defined by static variables 
  <code>yquotes.[.Daily, .Weekly, .Monthly, .Yearly]</code>
</p>

<pre>
  // Get historical prices for the last 3 years.
  prices, err := yquotes.HistoryForYears("AAPL", 3, yquotes.Daily)
  if err != nil {
    // handle error
  }
}
</pre>

<h4>Get historical prices between two dates</h4>
<p>
  Function <code>GetDailyHistory</code> accepts three arguments: symbol, date1 (from)
  date2 (to). Function returns list hisptorical prices <code>[]PriceH</code>. Dates are of time.Time type.
</p>

<pre>
  // Define layout of date. 
  layout := "2006-01-02"
  from := time.Parse(layout, "2012-01-01")
  to   := time.Now()

  prices, err := yquotes.GetDailyHistory("AAPL", from, to)
  if err != nil {

  }
</pre>

<h3>Data types</h3>
<h4>Stock type</h4>

<p>
  Notice that properies <code>Price</code> and <code>History</code> have different 
  types of price data. This is because historical data row has different data columns.
</p>

<pre>
  type Stock struct {
    // Symbol of stock that should meet requirements of Yahoo. Otherwise,
    // there will be no possibility to find stock.
    Symbol string `json:"symbol,omitempty"`

    // Name of the company will be filled from request of stock data.
    Name string `json:"name,omitempty"`

    // Information about last price of stock.
    Price *Price `json:"price,omitempty"`

    // Contains historical price information. If client asks information
    // for recent price, this field will be omited.
    History []PriceH `json:"history,omitempty"`
  }
</pre>

<h4>Price type</h4>

<p>Price struct represents price in single point in time.</p>
<pre>
  type Price struct {
    Bid           float64   `json:"bid,omitempty"`
    Ask           float64   `json:"ask,omitempty"`
    Open          float64   `json:"open,omitempty"`
    PreviousClose float64   `json:"previousClose,omitempty"`
    Last          float64   `json:"last,omitempty"`
    Date          time.Time `json:"date,omitempty"`
  }
</pre>

<h4>Historical price type</h4>

<p>This type represents row of historical price data.</p>
<pre>
  type PriceH struct {
    Date     time.Time `json:"date,omitempty"`
    Open     float64   `json:"open,omitempty"`
    High     float64   `json:"high,omitempty"`
    Low      float64   `json:"low,omitempty"`
    Close    float64   `json:"close,omitempty"`
    Volume   float64   `json:"volume,omitempty"`
    AdjClose float64   `json:"adjClose,omitempty"`
  }
</pre>

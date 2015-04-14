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
  // Get stock information without historical data.
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

<pre>
  // Price struct represents price in single point in time.
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

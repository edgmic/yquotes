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

<p>Load stock information without historical data.</p>

<p>
  <code>
    stock, err := yquotes.NewStock("AAPL", false)
    if err != nil {
      // handle error
    }

    symbol := stock.Symbol // AAPL
    name := stock.Name // Apple Inc.
    
    // Price information
    price := stock.Price 
    bid := price.Bid
    ask := price.Ask
    open := price.Open
    prevClose := price.PreviousClose
    last := price.Last
    date := price.Date 
  </code>
</p>
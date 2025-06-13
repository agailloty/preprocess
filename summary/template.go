package summary

var htmlTemplate string = `
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>CSV Summary Report</title>
  <style>
    body {
      font-family: sans-serif;
      margin: 2rem;
      background-color: #f7f7f7;
    }

    h1, h2 {
      color: #333;
    }

    .summary-box {
      background-color: #ffffff;
      border: 1px solid #ddd;
      padding: 1rem;
      margin-bottom: 2rem;
      border-radius: 6px;
    }

    table {
      border-collapse: collapse;
      width: 100%;
      background-color: #fff;
    }

    th, td {
      border: 1px solid #ccc;
      padding: 0.6rem;
      text-align: left;
    }

    th {
      background-color: #eaeaea;
    }

    .string-columns {
      margin-top: 2rem;
      display: grid;
      grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
      gap: 1rem;
    }

    .string-card {
      background-color: white;
      padding: 1rem;
      border-radius: 8px;
      border: 1px solid #ddd;
      box-shadow: 0 1px 3px rgba(0, 0, 0, 0.05);
    }

    .string-card h3 {
      margin: 0 0 0.5rem 0;
      font-size: 1.1rem;
      color: #333;
    }

    .string-card ul {
      padding-left: 1.2rem;
      margin: 0.5rem 0 0 0;
    }

    .section {
      margin-bottom: 2rem;
    }
  </style>
</head>
<body>

  <h1>CSV Summary Report</h1>

  <div class="summary-box">
    <h2>Dataset Overview</h2>
    <p><strong>Filename:</strong> {{.Data.Filename}}</p>
    <p><strong>Encoding:</strong> {{.Data.Encoding}}</p>
    <p><strong>CSV Separator:</strong> {{.Data.CsvSeparator}}</p>
    <p><strong>Decimal Separator:</strong> {{.Data.DecimalSeparator}}</p>
    <p><strong>Total Rows:</strong> {{.DataSummary.RowCount}}</p>
    <p><strong>Total Columns:</strong> {{.DataSummary.ColumnCount}}</p>
    <p><strong>Numeric Columns:</strong> {{.DataSummary.NumericColumns}}</p>
    <p><strong>String Columns:</strong> {{.DataSummary.StringColumns}}</p>
  </div>

  <div class="section">
    <h2>Numeric Columns</h2>
    <table>
      <thead>
        <tr>
          <th>Name</th>
          <th>Rows</th>
          <th>Missing</th>
          <th>Mean</th>
          <th>Median</th>
          <th>Min</th>
          <th>Max</th>
        </tr>
      </thead>
      <tbody>
        {{range .Columns}}
        {{if eq .Type "numeric"}}
        <tr>
          <td>{{.Name}}</td>
          <td>{{.RowCount}}</td>
          <td>{{.Missing}}</td>
          <td>{{printf "%.2f" .Mean}}</td>
          <td>{{printf "%.2f" .Median}}</td>
          <td>{{printf "%.2f" .Min}}</td>
          <td>{{printf "%.2f" .Max}}</td>
        </tr>
        {{end}}
        {{end}}
      </tbody>
    </table>
  </div>

  <div class="section">
    <h2>Categorical Columns</h2>
    <div class="string-columns">
      {{range .Columns}}
      {{if eq .Type "string"}}
      <div class="string-card">
        <h3>{{.Name}}</h3>
        <p><strong>Unique Values:</strong> {{.UniqueValueCount}}</p>
        {{if .UniqueValuesSummary}}
        <ul>
          {{range .UniqueValuesSummary}}
          <li>{{.Key}}: {{.Count}}</li>
          {{end}}
        </ul>
        {{else}}
        <p><em>No modalities available</em></p>
        {{end}}
      </div>
      {{end}}
      {{end}}
    </div>
  </div>

</body>
</html>

`

var diffTemplate string = `
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>CSV Diff Summary Report</title>
  <style>
    body {
      font-family: sans-serif;
      margin: 2rem;
      background-color: #f7f7f7;
    }

    h1, h2 {
      color: #333;
    }

    .summary-box {
      background-color: #ffffff;
      border: 1px solid #ddd;
      padding: 1rem;
      margin-bottom: 2rem;
      border-radius: 6px;
    }

    table {
      border-collapse: collapse;
      width: 100%;
      background-color: #fff;
    }

    th, td {
      border: 1px solid #ccc;
      padding: 0.6rem;
      text-align: left;
    }

    th {
      background-color: #eaeaea;
    }

    .string-columns {
      margin-top: 2rem;
      display: grid;
      grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
      gap: 1rem;
    }

    .string-card {
      background-color: white;
      padding: 1rem;
      border-radius: 8px;
      border: 1px solid #ddd;
      box-shadow: 0 1px 3px rgba(0, 0, 0, 0.05);
    }

    .string-card h3 {
      margin: 0 0 0.5rem 0;
      font-size: 1.1rem;
      color: #333;
    }

    .string-card ul {
      padding-left: 1.2rem;
      margin: 0.5rem 0 0 0;
    }

    .section {
      margin-bottom: 2rem;
    }

    .deleted {
      background-color: #f8d7da; /* rouge clair */
    }

    .added {
      background-color: #d4edda; /* vert clair */
    }

    .altered {
      background-color: #fff3cd; /* jaune moutarde clair */
    }
  </style>
</head>
<body>

  <h1>CSV Diff Summary Report</h1>

  <div class="summary-box">
    <h2>Dataset Overview</h2>
    <p><strong>Filename:</strong> {{.Data.Filename}}</p>
    <p><strong>Encoding:</strong> {{.Data.Encoding}}</p>
    <p><strong>CSV Separator:</strong> {{.Data.CsvSeparator}}</p>
    <p><strong>Decimal Separator:</strong> {{.Data.DecimalSeparator}}</p>
    <p><strong>Total Rows:</strong> {{.DataSummary.RowCount}}</p>
    <p><strong>Total Columns:</strong> {{.DataSummary.ColumnCount}}</p>
    <p><strong>Numeric Columns:</strong> {{.DataSummary.NumericColumns}}</p>
    <p><strong>String Columns:</strong> {{.DataSummary.StringColumns}}</p>
  </div>

  <div class="section">
    <h2>Numeric Columns</h2>
    <table>
      <thead>
        <tr>
          <th>Name</th>
          <th>Rows</th>
          <th>Missing</th>
          <th>Mean</th>
          <th>Median</th>
          <th>Min</th>
          <th>Max</th>
        </tr>
      </thead>
      <tbody>
        {{range .Columns}}
        {{if eq .Type "numeric"}}
        <tr class="{{if .IsDeleted}}deleted{{else if .IsAdded}}added{{else if .IsAltered}}altered{{end}}">
          <td>{{.Name}}</td>
          <td>{{.RowCount}}</td>
          <td>{{.Missing}}</td>
          <td>{{printf "%.2f" .Mean}}</td>
          <td>{{printf "%.2f" .Median}}</td>
          <td>{{printf "%.2f" .Min}}</td>
          <td>{{printf "%.2f" .Max}}</td>
        </tr>
        {{end}}
        {{end}}
      </tbody>
    </table>
  </div>

  <div class="section">
    <h2>Categorical Columns</h2>
    <div class="string-columns">
      {{range .Columns}}
      {{if eq .Type "string"}}
      <div class="string-card {{if .IsDeleted}}deleted{{else if .IsAdded}}added{{else if .IsAltered}}altered{{end}}">
        <h3>{{.Name}}</h3>
        <p><strong>Unique Values:</strong> {{.UniqueValueCount}}</p>

        {{if .AddedStringValues}}
        <p><strong>Added Values:</strong></p>
        <ul>
          {{range .AddedStringValues}}
          <li>+ {{.}}</li>
          {{end}}
        </ul>
        {{end}}

        {{if .RemovedStringValues}}
        <p><strong>Removed Values:</strong></p>
        <ul>
          {{range .RemovedStringValues}}
          <li>- {{.}}</li>
          {{end}}
        </ul>
        {{end}}

        {{if .UniqueValuesSummary}}
        <p><strong>Current Summary:</strong></p>
        <ul>
          {{range .UniqueValuesSummary}}
          <li>{{.Key}}: {{.Count}}</li>
          {{end}}
        </ul>
        {{else}}
        <p><em>No modalities available</em></p>
        {{end}}
      </div>
      {{end}}
      {{end}}
    </div>
  </div>

</body>
</html>
`

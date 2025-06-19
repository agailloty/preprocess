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

    .summary-box, .legend-box {
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

    .deleted { background-color: #f8d7da; }
    .added { background-color: #d4edda; }
    .altered-old { background-color: #fef3c7; }
    .altered-new { background-color: #e0f2fe; }

    .legend-box ul {
      list-style: none;
      padding-left: 0;
    }

    .legend-box li {
      display: inline-block;
      margin-right: 1rem;
    }

    .legend-box span {
      display: inline-block;
      width: 14px;
      height: 14px;
      margin-right: 5px;
      vertical-align: middle;
    }

    .color-added { background-color: #d4edda; }
    .color-deleted { background-color: #f8d7da; }
    .color-altered-old { background-color: #fef3c7; }
    .color-altered-new { background-color: #e0f2fe; }

    .two-col {
      display: flex;
      gap: 2rem;
    }

    .two-col > div {
      flex: 1;
    }
        tr.altered {
    background-color: #fff8dc;
  }

  .altered-old {
    background-color: #fef3c7;
    padding: 2px 4px;
    border-radius: 4px;
    display: inline-block;
  }

  .altered-new {
    background-color: #e0f2fe;
    padding: 2px 4px;
    border-radius: 4px;
    display: inline-block;
  }
  </style>
</head>
<body>

  <h1>CSV Diff Summary Report</h1>

  <div class="legend-box">
    <h2>Legend</h2>
    <ul>
      <li><span class="color-added"></span> Added</li>
      <li><span class="color-deleted"></span> Deleted</li>
      <li><span class="color-altered-old"></span> Altered (Before)</li>
      <li><span class="color-altered-new"></span> Altered (After)</li>
    </ul>
  </div>

  <div class="two-col">
    <div class="summary-box">
      <h2>Source Dataset</h2>
      <p><strong>Rows:</strong> {{.SourceDataSummary.RowCount}}</p>
      <p><strong>Columns:</strong> {{.SourceDataSummary.ColumnCount}}</p>
      <p><strong>Numeric Columns:</strong> {{.SourceDataSummary.NumericColumns}}</p>
      <p><strong>String Columns:</strong> {{.SourceDataSummary.StringColumns}}</p>
    </div>

    <div class="summary-box">
      <h2>Target Dataset</h2>
      <p><strong>Rows:</strong> {{.TargetDataSummary.RowCount}}</p>
      <p><strong>Columns:</strong> {{.TargetDataSummary.ColumnCount}}</p>
      <p><strong>Numeric Columns:</strong> {{.TargetDataSummary.NumericColumns}}</p>
      <p><strong>String Columns:</strong> {{.TargetDataSummary.StringColumns}}</p>
    </div>
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
      <tr class="{{if .IsDeleted}}deleted{{else if .IsAdded}}added{{else if .IsAltered}}.altered-old{{end}}">
        <td>{{.Name}}</td>
        <td>
          {{if .IsAltered}}
            <div class="altered-old">{{.OldStats.RowCount}}</div>
            <div class="altered-new">{{.NewStats.RowCount}}</div>
          {{else}}
            {{.RowCount}}
          {{end}}
        </td>
        <td>
          {{if .IsAltered}}
            <div class="altered-old">{{.OldStats.Missing}}</div>
            <div class="altered-new">{{.NewStats.Missing}}</div>
          {{else}}
            {{.Missing}}
          {{end}}
        </td>
        <td>
          {{if .IsAltered}}
            <div class="altered-old">{{printf "%.2f" .OldStats.Mean}}</div>
            <div class="altered-new">{{printf "%.2f" .NewStats.Mean}}</div>
          {{else}}
            {{printf "%.2f" .Mean}}
          {{end}}
        </td>
        <td>
          {{if .IsAltered}}
            <div class="altered-old">{{printf "%.2f" .OldStats.Median}}</div>
            <div class="altered-new">{{printf "%.2f" .NewStats.Median}}</div>
          {{else}}
            {{printf "%.2f" .Median}}
          {{end}}
        </td>
        <td>
          {{if .IsAltered}}
            <div class="altered-old">{{printf "%.2f" .OldStats.Min}}</div>
            <div class="altered-new">{{printf "%.2f" .NewStats.Min}}</div>
          {{else}}
            {{printf "%.2f" .Min}}
          {{end}}
        </td>
        <td>
          {{if .IsAltered}}
            <div class="altered-old">{{printf "%.2f" .OldStats.Max}}</div>
            <div class="altered-new">{{printf "%.2f" .NewStats.Max}}</div>
          {{else}}
            {{printf "%.2f" .Max}}
          {{end}}
        </td>
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
      <div class="string-card {{if .IsDeleted}}deleted{{else if .IsAdded}}added{{else if .IsAltered}}altered-old{{end}}">
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

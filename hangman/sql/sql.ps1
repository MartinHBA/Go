# Define the function
function Get-SQLServerCUList {

    # Declare parameters
    param (
        [Parameter(Mandatory = $true)]
        [ValidateSet(2008, 2008R2, 2012, 2014, 2016)]
        [int]$MajorVersion
    )

    # Define a hashtable of URLs for each major version
    $URLs = @{
        2008   = "https://support.microsoft.com/en-us/help/957826/the-builds-for-all-sql-server-versions"
        2008R2 = "https://support.microsoft.com/en-us/help/2567616/the-builds-for-all-sql-server-versions"
        2012   = "https://support.microsoft.com/en-us/help/2692828/the-builds-for-all-sql-server-versions"
        2014   = "https://support.microsoft.com/en-us/help/2936603/the-builds-for-all-sql-server-versions"
        2016   = "https://support.microsoft.com/en-us/help/3177312/the-builds-for-all-sql-server-versions"
    }

    # Get the URL for the major version
    $URL = $URLs[$MajorVersion]

    # Invoke a web request to get the HTML content of the URL
    $HTML = Invoke-WebRequest -Uri $URL

    # Define a regular expression pattern to match the cumulative update information in the HTML content
    $Pattern = "(?smi)<tr>\s*<td>\s*(?<MajorVersion>\d+)\.(?<MinorVersion>\d+)\.(?<Build>\d+)\.0\s*</td>\s*<td>(?<ReleaseDate>.*?)</td>\s*<td>(?<DownloadLink>.*?)</td>"

    # Use the Select-Xml cmdlet to convert the HTML content to XML and then use XPath to select only the table rows that contain cumulative update information
    $XMLRows = Select-Xml -Content $HTML.Content -XPath "//table[@class='b-hgrid']/tbody/tr[position() > 1]"

    # Initialize an empty array to store the cumulative update objects
    $CUList = @()

    # Loop through each XML row and parse it using the regular expression pattern
    foreach ($XMLRow in $XMLRows) {

        # Convert the XML row to a string and match it with the pattern
        $Match = [regex]::Match($XMLRow.Node.OuterXml, $Pattern)

        # If there is a match, create a custom object with properties from the named capture groups and add it to the array
        if ($Match.Success) {

            # Create a custom object with properties from the named capture groups
            $CUObject = New-Object -TypeName PSObject -Property @{
                MajorVersion = [int]$Match.Groups["MajorVersion"].Value
                MinorVersion = [int]$Match.Groups["MinorVersion"].Value
                Build        = [int]$Match.Groups["Build"].Value
                ReleaseDate  = [datetime]$Match.Groups["ReleaseDate"].Value.Trim()
                DownloadLink = ($Match.Groups["DownloadLink"].Value.Trim() -replace "<.*?>", "").Trim()
            }

            # Add the custom object to the array
            $CUList += $CUObject

        }

    }

    # Return the array of cumulative update objects
    return $CUList

}

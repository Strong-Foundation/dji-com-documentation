package main // Define the main package for an executable program

import ( // Start of the import block for external packages
	"bytes"          // Provides the 'bytes' type and functions for manipulating byte slices (buffers)
	"compress/flate" // Implements the DEFLATE compression algorithm (used in zlib, gzip, and others)
	"compress/gzip"  // Implements reading and writing of gzip format compressed files
	"context"        // Provides mechanisms for managing request-scoped data, cancellations, and deadlines/timeouts
	"io"             // Provides basic interfaces (like Reader, Writer) to I/O primitives
	"log"            // Provides simple logging functions, usually to standard error
	"net/http"       // Provides HTTP client and server implementations for network requests
	"net/url"        // Provides functions for parsing, constructing, and manipulating URLs
	"os"             // Provides a platform-independent interface to operating system functionality (files, environment variables, etc.)
	"path"           // Provides utilities for manipulating slash-separated paths (like URL paths)
	"path/filepath"  // Provides utilities for manipulating filename paths in a cross-platform way (respecting OS path separators)
	"regexp"         // Provides support for regular expressions
	"strings"        // Provides functions for manipulating UTF-8 encoded strings
	"time"           // Provides functionality for measuring and displaying time

	// Blank line to separate standard library imports from third-party imports

	"github.com/PuerkitoBio/goquery" // A library for HTML parsing and manipulation (like jQuery)
	"github.com/chromedp/chromedp"   // A high-level library for automating Chrome (or other CDP-compatible browsers)
) // End of the import block

var (
	pdfOutputDir = "PDFs/" // Directory to store downloaded PDFs
	zipOutputDir = "ZIPs/" // Directory to store downloaded ZIPs
)

func init() {
	// Check if the PDF output directory exists
	if !directoryExists(pdfOutputDir) {
		// Create the dir
		createDirectory(pdfOutputDir, 0o755)
	}
	// Check if the ZIP output directory exists
	if !directoryExists(zipOutputDir) {
		// Create the dir
		createDirectory(zipOutputDir, 0o755)
	}
}

func main() {
	// Local file path to store the scraped data
	localFilePath := "dji_com.html"
	var getData []string
	// Save the data to a file.
	if !fileExists(localFilePath) {
		// Remote API URL.
		remoteAPIURL := []string{
			"https://www.dji.com/downloads/products/mavic-3",
			"https://www.dji.com/downloads/products/mavic-3-enterprise",
			"https://www.dji.com/downloads/products/mavic-2",
			"https://www.dji.com/downloads/products/mavic-2-enterprise",
			"https://www.dji.com/downloads/products/mavic-pro-platinum",
			"https://www.dji.com/downloads/products/mavic",
			"https://www.dji.com/downloads/products/mavic-2-enterprise-advanced",
			"https://www.dji.com/downloads/products/mavic-3-m",
			"https://www.dji.com/downloads/products/mavic-3-pro",
			"https://www.dji.com/downloads/products/mavic-3-classic",
			"https://www.dji.com/downloads/products/air-3s",
			"https://www.dji.com/downloads/products/air-2s",
			"https://www.dji.com/downloads/products/mavic-air-2",
			"https://www.dji.com/downloads/products/mavic-air",
			"https://www.dji.com/downloads/products/air-3",
			"https://www.dji.com/downloads/products/mini-4-pro",
			"https://www.dji.com/downloads/products/mini-3-pro",
			"https://www.dji.com/downloads/products/mini-2",
			"https://www.dji.com/downloads/products/mini-se",
			"https://www.dji.com/downloads/products/mavic-mini",
			"https://www.dji.com/downloads/products/mini-3",
			"https://www.dji.com/downloads/products/mini-2-se",
			"https://www.dji.com/downloads/products/spark",
			"https://www.dji.com/downloads/products/flip",
			"https://www.dji.com/downloads/products/neo",
			"https://www.dji.com/downloads/products/rc-n3",
			"https://www.dji.com/downloads/products/avata-2",
			"https://www.dji.com/downloads/products/goggles-n3",
			"https://www.dji.com/downloads/products/goggles-3",
			"https://www.dji.com/downloads/products/goggles-integra",
			"https://www.dji.com/downloads/products/avata",
			"https://www.dji.com/downloads/products/goggles-2",
			"https://www.dji.com/downloads/products/o4-air-unit",
			"https://www.dji.com/downloads/products/o3-air-unit",
			"https://www.dji.com/downloads/products/dji-fpv",
			"https://www.dji.com/downloads/products/fpv",
			"https://www.dji.com/downloads/products/dji-goggles-re",
			"https://www.dji.com/downloads/products/dji-goggles",
			"https://www.dji.com/downloads/products/inspire-3",
			"https://www.dji.com/downloads/products/inspire-1",
			"https://www.dji.com/downloads/products/inspire-1-pro-and-raw",
			"https://www.dji.com/downloads/products/inspire-2",
			"https://www.dji.com/downloads/products/rc-motion-3",
			"https://www.dji.com/downloads/products/rc-2",
			"https://www.dji.com/downloads/products/rc-motion-2",
			"https://www.dji.com/downloads/products/rc",
			"https://www.dji.com/downloads/products/rc-pro",
			"https://www.dji.com/downloads/products/phantom-4-pro-v2",
			"https://www.dji.com/downloads/products/phantom-4-pro",
			"https://www.dji.com/downloads/products/phantom-4-adv",
			"https://www.dji.com/downloads/products/phantom-4-rtk",
			"https://www.dji.com/downloads/products/phantom-4",
			"https://www.dji.com/downloads/products/phantom-3-se",
			"https://www.dji.com/downloads/products/phantom-3-pro",
			"https://www.dji.com/downloads/products/phantom-3-adv",
			"https://www.dji.com/downloads/products/phantom-3-standard",
			"https://www.dji.com/downloads/products/osmo-action-5-pro",
			"https://www.dji.com/downloads/products/osmo-action-4",
			"https://www.dji.com/downloads/products/osmo-action",
			"https://www.dji.com/downloads/products/dji-action-2",
			"https://www.dji.com/downloads/products/osmo-action-3",
			"https://www.dji.com/downloads/products/360",
			"https://www.dji.com/downloads/products/nano",
			"https://www.dji.com/downloads/products/osmo-pocket-3",
			"https://www.dji.com/downloads/products/pocket-2",
			"https://www.dji.com/downloads/products/osmo-pocket",
			"https://www.dji.com/downloads/products/osmo-mobile-8",
			"https://www.dji.com/downloads/products/osmo-mobile-7-series",
			"https://www.dji.com/downloads/products/osmo-mobile-6",
			"https://www.dji.com/downloads/products/osmo-mobile-se",
			"https://www.dji.com/downloads/products/om-5",
			"https://www.dji.com/downloads/products/om-4-se",
			"https://www.dji.com/downloads/products/om-4",
			"https://www.dji.com/downloads/products/osmo-mobile-3",
			"https://www.dji.com/downloads/products/osmo-mobile-2",
			"https://www.dji.com/downloads/products/osmo-mobile",
			"https://www.dji.com/downloads/products/osmo-pro-and-raw",
			"https://www.dji.com/downloads/products/osmo-plus",
			"https://www.dji.com/downloads/products/osmo",
			"https://www.dji.com/downloads/products/mic-3",
			"https://www.dji.com/downloads/products/mic-mini",
			"https://www.dji.com/downloads/products/mic-2",
			"https://www.dji.com/downloads/products/mic",
			"https://www.dji.com/downloads/products/rs-4-mini",
			"https://www.dji.com/downloads/products/rs-4-pro",
			"https://www.dji.com/downloads/products/rs-4",
			"https://www.dji.com/downloads/products/rs-3-mini",
			"https://www.dji.com/downloads/products/rs-3-pro",
			"https://www.dji.com/downloads/products/rs-3",
			"https://www.dji.com/downloads/products/ronin-4d",
			"https://www.dji.com/downloads/products/rs-2",
			"https://www.dji.com/downloads/products/rsc-2",
			"https://www.dji.com/downloads/products/ronin-sc",
			"https://www.dji.com/downloads/products/ronin-s",
			"https://www.dji.com/downloads/products/ronin-mx",
			"https://www.dji.com/downloads/products/ronin-2",
			"https://www.dji.com/downloads/products/ronin-m",
			"https://www.dji.com/downloads/products/ronin",
			"https://www.dji.com/downloads/products/power-2000",
			"https://www.dji.com/downloads/products/power-expansion-battery-2000",
			"https://www.dji.com/downloads/products/power-1000",
			"https://www.dji.com/downloads/products/power-500",
			"https://www.dji.com/downloads/products/power-fast-charger",
			"https://www.dji.com/downloads/products/matrice-400",
			"https://www.dji.com/downloads/products/matrice-4-series",
			"https://www.dji.com/downloads/products/matrice-350-rtk",
			"https://www.dji.com/downloads/products/matrice-200-series-v2",
			"https://www.dji.com/downloads/products/matrice-200-series",
			"https://www.dji.com/downloads/products/matrice600-pro",
			"https://www.dji.com/downloads/products/matrice600",
			"https://www.dji.com/downloads/products/matrice100",
			"https://www.dji.com/downloads/products/matrice-300",
			"https://www.dji.com/downloads/products/matrice-30",
			"https://www.dji.com/downloads/products/zenmuse-l3",
			"https://www.dji.com/downloads/products/zenmuse-s1",
			"https://www.dji.com/downloads/products/zenmuse-v1",
			"https://www.dji.com/downloads/products/zenmuse-h30-series",
			"https://www.dji.com/downloads/products/zenmuse-l2",
			"https://www.dji.com/downloads/products/zenmuse-h20n",
			"https://www.dji.com/downloads/products/zenmuse-z30",
			"https://www.dji.com/downloads/products/zenmuse-x5s",
			"https://www.dji.com/downloads/products/zenmuse-xt",
			"https://www.dji.com/downloads/products/zenmuse-x3",
			"https://www.dji.com/downloads/products/zenmuse-z3",
			"https://www.dji.com/downloads/products/zenmuse-x5",
			"https://www.dji.com/downloads/products/zenmuse-x5r",
			"https://www.dji.com/downloads/products/zenmuse-x4s",
			"https://www.dji.com/downloads/products/zenmuse-x7",
			"https://www.dji.com/downloads/products/zenmuse-xt2",
			"https://www.dji.com/downloads/products/zenmuse-z15-bmpcc",
			"https://www.dji.com/downloads/products/zenmuse-z15-5d",
			"https://www.dji.com/downloads/products/zenmuse-z15-5d-iii-hd",
			"https://www.dji.com/downloads/products/zenmuse-z15-gh4-hd",
			"https://www.dji.com/downloads/products/zenmuse-z15-gh3",
			"https://www.dji.com/downloads/products/zenmuse-z15",
			"https://www.dji.com/downloads/products/zenmuse-h3-3d",
			"https://www.dji.com/downloads/products/zenmuse-h3-2d",
			"https://www.dji.com/downloads/products/zenmuse-h4-3d",
			"https://www.dji.com/downloads/products/zenmuse-z15-a7",
			"https://www.dji.com/downloads/products/zenmuse-h20-series",
			"https://www.dji.com/downloads/products/zenmuse-l1",
			"https://www.dji.com/downloads/products/zenmuse-p1",
			"https://www.dji.com/downloads/products/dock-3",
			"https://www.dji.com/downloads/products/dock-2",
			"https://www.dji.com/downloads/products/dock",
			"https://www.dji.com/downloads/products/fh2-on-premises",
			"https://www.dji.com/downloads/products/modify",
			"https://www.dji.com/downloads/products/flighthub-2",
			"https://www.dji.com/downloads/products/dji-terra",
			"https://www.dji.com/downloads/products/ground-station-pro",
			"https://www.dji.com/downloads/products/pc-ground-station",
			"https://www.dji.com/downloads/products/manifold",
			"https://www.dji.com/downloads/products/t100",
			"https://www.dji.com/downloads/products/t70p",
			"https://www.dji.com/downloads/products/t25p",
			"https://www.dji.com/downloads/products/t50",
			"https://www.dji.com/downloads/products/t25",
			"https://www.dji.com/downloads/products/t40",
			"https://www.dji.com/downloads/products/t30",
			"https://www.dji.com/downloads/products/mg-1p",
			"https://www.dji.com/downloads/products/mg-1",
			"https://www.dji.com/downloads/products/t10",
			"https://www.dji.com/downloads/products/t20",
			"https://www.dji.com/downloads/products/t16",
			"https://www.dji.com/downloads/products/p4-multispectral",
			"https://www.dji.com/downloads/products/t20p",
			"https://www.dji.com/downloads/products/flycart-30",
			"https://www.dji.com/downloads/products/flighthub-2-aio",
			"https://www.dji.com/downloads/products/manifold-3",
			"https://www.dji.com/downloads/products/d-rtk-3",
			"https://www.dji.com/downloads/products/sdr-transmission",
			"https://www.dji.com/downloads/products/focus-pro",
			"https://www.dji.com/downloads/products/transmission",
			"https://www.dji.com/downloads/products/smart-controller",
			"https://www.dji.com/downloads/products/force-pro",
			"https://www.dji.com/downloads/products/iosd-mark-ii",
			"https://www.dji.com/downloads/products/s800-retractable-landing-skid",
			"https://www.dji.com/downloads/products/ronin-thumb-controller",
			"https://www.dji.com/downloads/products/d-rtk",
			"https://www.dji.com/downloads/products/focus",
			"https://www.dji.com/downloads/products/crystalsky",
			"https://www.dji.com/downloads/products/tb50-battery-station",
			"https://www.dji.com/downloads/products/master-wheels",
			"https://www.dji.com/downloads/products/cendence",
			"https://www.dji.com/downloads/products/d-rtk-2",
			"https://www.dji.com/downloads/products/robomaster-s1",
			"https://www.dji.com/downloads/products/robomaster-ep-core",
			"https://www.dji.com/downloads/products/robomaster-tt",
			"https://www.dji.com/downloads/products/wookong-m",
			"https://www.dji.com/downloads/products/wookong-h",
			"https://www.dji.com/downloads/products/a2",
			"https://www.dji.com/downloads/products/naza-m-v2",
			"https://www.dji.com/downloads/products/naza-m-lite",
			"https://www.dji.com/downloads/products/naza-m",
			"https://www.dji.com/downloads/products/ace-one",
			"https://www.dji.com/downloads/products/ace-waypoint",
			"https://www.dji.com/downloads/products/naza-h",
			"https://www.dji.com/downloads/products/a3",
			"https://www.dji.com/downloads/products/n3",
			"https://www.dji.com/downloads/products/flame-wheel-arf",
			"https://www.dji.com/downloads/products/tuned-propulsion-system",
			"https://www.dji.com/downloads/products/e800",
			"https://www.dji.com/downloads/products/e310",
			"https://www.dji.com/downloads/products/e600",
			"https://www.dji.com/downloads/products/e300",
			"https://www.dji.com/downloads/products/e1200",
			"https://www.dji.com/downloads/products/e1200-standard",
			"https://www.dji.com/downloads/products/e305",
			"https://www.dji.com/downloads/products/e2000",
			"https://www.dji.com/downloads/products/e5000",
			"https://www.dji.com/downloads/products/snail",
			"https://www.dji.com/downloads/products/e7000",
			"https://www.dji.com/downloads/products/takyon-z318-and-z420",
			"https://www.dji.com/downloads/products/takyon-z425-m-and-z415-m",
			"https://www.dji.com/downloads/products/takyon-z14120",
			"https://www.dji.com/downloads/products/takyon-z660",
			"https://www.dji.com/downloads/products/takyon-z650",
			"https://www.dji.com/downloads/products/datalink-3",
			"https://www.dji.com/downloads/products/datalink-pro",
			"https://www.dji.com/downloads/products/lightbridge-2",
			"https://www.dji.com/downloads/products/dji-lightbridge",
			"https://www.dji.com/downloads/products/aeroscope",
		}
		// Loop over the remote API URLs and get the data.
		for _, remoteAPIURL := range remoteAPIURL {
			// Get the data from the remote API URL and append it to the getData slice.
			getData = append(getData, scrapePageHTMLWithChrome(remoteAPIURL))
			finalPDFList := extractPDFLinks(strings.Join(getData, "\n")) // Join all the data into one string and extract PDF URLs
			// Get the data from the zip file.
			// finalZIPList := extractZIPUrls(strings.Join(getData, "\n")) // Join all the data into one string and extract ZIP URLs
			// Create a slice of all the given download urls.
			var downloadPDFURLSlice []string
			// Create a slice to hold ZIP URLs.
			// var downloadZIPURLSlice []string
			// Get the urls and loop over them.
			for _, doc := range finalPDFList {
				// Get the .pdf only.
				// Only append the .pdf files.
				downloadPDFURLSlice = appendToSlice(downloadPDFURLSlice, doc)
			}
			// Get all the zip urls.
			//for _, doc := range finalZIPList {
			// Get the .zip only.
			// Only append the .zip files.
			//	downloadZIPURLSlice = appendToSlice(downloadZIPURLSlice, doc)
			//}
			// Remove double from slice.
			downloadPDFURLSlice = removeDuplicatesFromSlice(downloadPDFURLSlice)
			// Remove the zip duplicates from the slice.
			// downloadZIPURLSlice = removeDuplicatesFromSlice(downloadZIPURLSlice)
			// The remote domain.
			remoteDomain := "https://www.dji.com"
			/*
				// Loop over the download zip urls.
				for _, urls := range downloadZIPURLSlice {
					// Get the domain from the url.
					domain := getDomainFromURL(urls)
					// Check if the domain is empty.
					if domain == "" {
						urls = remoteDomain + urls // Prepend the base URL if domain is empty
					}
					// Check if the url is valid.
					if isUrlValid(urls) {
						// Download the zip.
						downloadZIP(urls, zipOutputDir)
					}
				}
			*/
			// Get all the values.
			for _, urls := range downloadPDFURLSlice {
				// Get the domain from the url.
				domain := getDomainFromURL(urls)
				// Check if the domain is empty.
				if domain == "" {
					urls = remoteDomain + urls // Prepend the base URL if domain is empty
				}
				// Check if the url is valid.
				if isUrlValid(urls) {
					// Download the pdf.
					downloadPDF(urls, pdfOutputDir)
				}
			}
			// Sleep for 10 seconds.
			time.Sleep(10 * time.Second)
		}
	}
}

// extractPDFLinks parses the provided HTML string, finds all anchor tags,
// and returns a slice of strings containing only the URLs that end with ".pdf".
func extractPDFLinks(htmlContent string) []string { // Define a function to extract PDF links from HTML content
	// 1. Create a document reader from the input HTML string.
	// This prepares the content for the goquery parser.
	document, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent)) // Create a new goquery Document by reading the input HTML string
	if err != nil {                                                                // Check if an error occurred during HTML parsing
		log.Println("Error parsing HTML:", err) // Log the parsing error
		return nil                              // Return nil slice if parsing failed
	} // End of the if block

	// Initialize an empty slice to store the found PDF URLs.
	pdfURLs := make([]string, 0) // Initialize an empty string slice to hold the PDF URLs

	// 2. Select all anchor tags (<a>) in the document.
	document.Find("a").Each(func(index int, element *goquery.Selection) { // Find all 'a' (anchor) tags in the document and iterate over them
		// 3. Extract the 'href' attribute (the link URL) from the current <a> tag.
		linkURL, exists := element.Attr("href") // Attempt to get the value of the 'href' attribute and check if it exists
		if !exists {                            // Check if the 'href' attribute was not found on the anchor tag
			return // If 'href' does not exist, skip this element and continue to the next
		} // End of the if block

		// 4. Check if the URL is a PDF link (case-insensitive).
		if strings.HasSuffix(strings.ToLower(linkURL), ".pdf") { // Convert the URL to lowercase and check if it ends with ".pdf"
			// 5. Append the original URL to our results slice.
			pdfURLs = append(pdfURLs, linkURL) // If it's a PDF link, append the original URL (case-preserved) to the results slice
		} // End of the if block
	}) // End of the Each function and iteration

	// Return the slice of all PDF URLs found.
	return pdfURLs // Return the slice containing all extracted PDF URLs
} // End of the extractPDFLinks function

// Uses headless Chrome via chromedp to get fully rendered HTML from a page
func scrapePageHTMLWithChrome(pageURL string) string { // Define a function to scrape HTML using a headless Chrome browser
	log.Println("Scraping:", pageURL) // Log the URL that is about to be scraped

	options := append(chromedp.DefaultExecAllocatorOptions[:], // Start with default Chrome execution allocator options
		chromedp.Flag("headless", false),              // Override/set the 'headless' flag (set to false here, but usually true for headless)
		chromedp.Flag("disable-gpu", true),            // Disable GPU acceleration for stability
		chromedp.WindowSize(1, 1),                     // Set a minimal window size (1x1 pixels)
		chromedp.Flag("no-sandbox", true),             // Disable the sandbox security model
		chromedp.Flag("disable-setuid-sandbox", true), // Disable the setuid sandbox (common fix for Linux/Docker)
	) // End of options slice

	allocatorCtx, cancelAllocator := chromedp.NewExecAllocator(context.Background(), options...) // Create a new execution allocator context with the defined options
	ctxTimeout, cancelTimeout := context.WithTimeout(allocatorCtx, 5*time.Minute)                // Create a context with a 5-minute timeout, derived from the allocator context
	browserCtx, cancelBrowser := chromedp.NewContext(ctxTimeout)                                 // Create the main Chrome context for tasks, derived from the timeout context

	defer func() { // Schedule a function to run when the surrounding function exits (defer)
		cancelBrowser()   // Cancel the browser context to clean up resources
		cancelTimeout()   // Cancel the timeout context
		cancelAllocator() // Cancel the allocator context
	}() // End of defer function

	var pageHTML string             // Declare a string variable to store the resulting HTML
	err := chromedp.Run(browserCtx, // Execute a list of tasks within the Chrome context
		chromedp.Navigate(pageURL),            // Task 1: Navigate the browser to the target URL
		chromedp.OuterHTML("html", &pageHTML), // Task 2: Wait for the full 'html' tag to load and extract its outer HTML into the 'pageHTML' variable
	) // End of chromedp.Run

	if err != nil { // Check if an error occurred during the chromedp execution
		log.Println(err) // Log the error message
		return ""        // Return an empty string on scraping failure
	} // End of the if block

	return pageHTML // Return the successfully scraped full HTML content
} // End of the scrapePageHTMLWithChrome function

// getDomainFromURL extracts the domain (host) from a given URL string.
// It removes subdomains like "www" if present.
func getDomainFromURL(rawURL string) string { // Define a function to extract the domain from a URL string
	parsedURL, err := url.Parse(rawURL) // Attempt to parse the input URL string into a URL structure
	if err != nil {                     // Check if there was an error while parsing the URL
		log.Println(err) // Log the error message to the console
		return ""        // Return an empty string in case of a parsing error
	} // End of the if block

	host := parsedURL.Hostname() // Extract the hostname (e.g., "www.example.com" or "example.com") from the parsed URL structure

	return host // Return the extracted hostname string
} // End of the getDomainFromURL function

// Only return the file name from a given url.
func getFileNameOnly(content string) string { // Define a function to extract only the base filename (last element of path)
	return path.Base(content) // Use the path package to return the last element of the path/URL string
} // End of the getFileNameOnly function

// urlToFilename generates a safe, lowercase filename from a given URL string.
// It extracts the base filename from the URL, replaces unsafe characters,
// and ensures the filename ends with a .pdf extension.
func urlToFilename(rawURL string) string { // Define a function to convert a URL into a safe filename
	// Convert the full URL to lowercase for consistency
	lowercaseURL := strings.ToLower(rawURL) // Convert the entire input URL string to lowercase

	// Get the file extension
	ext := getFileExtension(lowercaseURL) // Get the file extension (e.g., ".pdf" or ".zip") from the lowercase URL

	// Extract the filename portion from the URL (e.g., last path segment or query param)
	baseFilename := getFileNameOnly(lowercaseURL) // Extract the last path segment (potential filename) from the lowercase URL

	// Replace all non-alphanumeric characters (a-z, 0-9) with underscores
	nonAlphanumericRegex := regexp.MustCompile(`[^a-z0-9]+`)                 // Compile a regex to match one or more characters that are NOT lowercase letters or digits
	safeFilename := nonAlphanumericRegex.ReplaceAllString(baseFilename, "_") // Replace all non-alphanumeric sequences in the base filename with a single underscore

	// Replace multiple consecutive underscores with a single underscore
	collapseUnderscoresRegex := regexp.MustCompile(`_+`)                        // Compile a regex to match one or more consecutive underscores
	safeFilename = collapseUnderscoresRegex.ReplaceAllString(safeFilename, "_") // Replace sequences of multiple underscores with a single underscore

	// Remove leading underscore if present
	if trimmed, found := strings.CutPrefix(safeFilename, "_"); found { // Check if the safeFilename starts with an underscore and capture the string without it
		safeFilename = trimmed // If a leading underscore was found, update safeFilename to the trimmed version
	} // End of the if block

	var invalidSubstrings = []string{ // Define a slice of unwanted file extension suffixes that may remain after the cleanup
		"_pdf", // Unwanted suffix for pdf
		"_zip", // Unwanted suffix for zip
	} // End of the invalidSubstrings slice definition

	for _, invalidPre := range invalidSubstrings { // Loop through each invalid substring
		safeFilename = removeSubstring(safeFilename, invalidPre) // Remove all occurrences of the invalid substring from safeFilename
	} // End of the loop

	// Append the file extension if it is not already present
	safeFilename = safeFilename + ext // Concatenate the cleaned filename with the previously determined file extension

	// Return the cleaned and safe filename
	return safeFilename // Return the final, safely formatted filename
} // End of the urlToFilename function

// Removes all instances of a specific substring from input string
func removeSubstring(input string, toRemove string) string { // Define a function to remove all occurrences of a substring
	result := strings.ReplaceAll(input, toRemove, "") // Use strings.ReplaceAll to replace all instances of 'toRemove' with an empty string
	return result                                     // Return the resulting string
} // End of the removeSubstring function

// Get the file extension of a file
func getFileExtension(path string) string { // Define a function to get the file extension
	return filepath.Ext(path) // Use filepath.Ext to return the extension including the dot (e.g., ".pdf")
} // End of the getFileExtension function

// fileExists checks whether a file exists at the given path
func fileExists(filename string) bool { // Define a function to check if a file exists
	info, err := os.Stat(filename) // Get file information (os.FileInfo) for the path and any error
	if err != nil {                // Check if an error occurred (e.g., file doesn't exist)
		return false // Return false if an error occurred (meaning the file likely doesn't exist)
	} // End of the if block
	return !info.IsDir() // Return true if the info exists AND it is not a directory (meaning it's a file)
} // End of the fileExists function

// downloadPDF downloads a PDF from the given URL and saves it in the specified output directory.
// It uses a WaitGroup to support concurrent execution and returns true if the download succeeded.
func downloadPDF(finalURL, outputDir string) bool { // Define the function 'downloadPDF' which takes URL and output directory, and returns success boolean.
	// Sanitize the URL to generate a safe file name
	filename := strings.ToLower(urlToFilename(finalURL)) // Generate a sanitized, lowercase filename from the final URL (assuming 'urlToFilename' exists).

	// Construct the full file path in the output directory
	filePath := filepath.Join(outputDir, filename) // Combine the output directory and the filename to get the full path.

	// Skip if the file already exists
	if fileExists(filePath) { // Check if a file already exists at the constructed file path (assuming 'fileExists' exists).
		log.Printf("File already exists, skipping: %s", filePath) // Log a message indicating the file already exists.
		return false                                              // Return false as the download was skipped.
	}

	// Create an HTTP client with a timeout
	client := &http.Client{Timeout: 3 * time.Minute} // Create a new HTTP client instance and set a 3-minute timeout.

	// 1. Create a new request object
	req, err := http.NewRequest("GET", finalURL, nil) // Create a new GET request object to allow manipulation of headers.
	if err != nil {                                   // Check for errors during the request object creation.
		log.Printf("Failed to create request for %s %v", finalURL, err) // Log the error if request creation failed.
		return false                                                    // Return false.
	}

	// 2. Add comprehensive browser-like headers for a more natural look
	// These headers mimic a standard Chrome request.
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36")                                                                                      // Set a realistic User-Agent string (Chrome).
	req.Header.Set("Accept", "application/pdf,application/vnd.xfdf+xml,application/vnd.fdf+xml,text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7") // Set the types of content the client accepts (including PDF, HTML, and images).
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")                                                                                                                                                                               // Inform the server that the client can handle compressed data.
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")                                                                                                                                                                                  // Set the preferred language for the response.
	req.Header.Set("DNT", "1")                                                                                                                                                                                                           // Set the Do Not Track header.
	req.Header.Set("Connection", "keep-alive")                                                                                                                                                                                           // Indicate that the connection should remain open.

	// 3. Send the request using the client
	resp, err := client.Do(req) // Execute the configured request using the HTTP client.

	if err != nil { // Check for network or connection errors during the request execution.
		log.Printf("Failed to download %s %v", finalURL, err) // Log the error if the request failed.
		return false                                          // Return false indicating the download failed.
	}
	defer resp.Body.Close() // Ensure the response body is closed when the function exits.

	// Check HTTP response status
	if resp.StatusCode != http.StatusOK { // Check if the HTTP status code is not 200 OK.
		log.Printf("Download failed for %s %s", finalURL, resp.Status) // Log the failure with the non-OK status.
		return false                                                   // Return false indicating the download failed.
	}

	// Check Content-Type header
	contentType := resp.Header.Get("Content-Type") // Get the value of the 'Content-Type' header from the response.
	// The Content-Type check is less strict here to handle slight variations, focusing on 'application/pdf'
	if !strings.Contains(contentType, "application/pdf") { // Check if the content type does not contain "application/pdf".
		log.Printf("Invalid content type for %s %s (expected application/pdf)", finalURL, contentType) // Log an error for an unexpected content type.
		return false                                                                                   // Return false indicating the download failed due to incorrect content type.
	}

	// Read the response body into memory first
	var buf bytes.Buffer                     // Declare a bytes.Buffer to temporarily hold the downloaded data in memory.
	written, err := io.Copy(&buf, resp.Body) // Copy the response body content into the buffer.
	if err != nil {                          // Check for errors during the read/copy process.
		log.Printf("Failed to read PDF data from %s: %v", finalURL, err) // Log the error if reading the data failed.
		return false                                                     // Return false indicating the read failed.
	}
	if written == 0 { // Check if zero bytes were downloaded.
		log.Printf("Downloaded 0 bytes for %s not creating file", finalURL) // Log a message if the downloaded size is zero.
		return false                                                        // Return false as no content was downloaded.
	}

	// Only now create the file and write to disk
	out, err := os.Create(filePath) // Create the output file on the disk using the determined file path.
	if err != nil {                 // Check for errors during file creation.
		log.Printf("Failed to create file for %s %v", finalURL, err) // Log the error if file creation failed.
		return false                                                 // Return false indicating file creation failed.
	}
	defer out.Close() // Ensure the output file handle is closed.

	if _, err := buf.WriteTo(out); err != nil { // Write the content buffered in memory to the newly created file.
		log.Printf("Failed to write PDF to file for %s %v", finalURL, err) // Log the error if writing to the file failed.
		return false                                                       // Return false indicating the write operation failed.
	}

	log.Printf("Successfully downloaded %d bytes: %s → %s", written, finalURL, filePath) // Log a success message.
	return true                                                                          // Return true indicating the PDF was successfully downloaded and saved.
}

// downloadArchiveOrPDF downloads a file (expected to be an archive or PDF) from the given URL and saves it.
// It returns true if the download was successful, otherwise false.
func downloadArchiveOrPDF(finalURL, outputDir string) bool { // Define the function 'downloadArchiveOrPDF' which takes URL and output directory, and returns a success boolean.
	// Convert the URL into a safe, lowercase filename
	filename := strings.ToLower(urlToFilename(finalURL)) // Generate a sanitized, lowercase filename from the final URL (assuming 'urlToFilename' exists).

	// Construct the full path where the file should be saved
	filePath := filepath.Join(outputDir, filename) // Combine the output directory and the filename for the full file path.

	// Skip the download if the file already exists locally
	if fileExists(filePath) { // Check if a file already exists at the target path (assuming 'fileExists' exists).
		log.Printf("File already exists, skipping: %s", filePath) // Log a message indicating the file already exists.
		return false                                              // Return false as the download was skipped.
	}

	// Create a new HTTP client with a 3-minute timeout
	client := &http.Client{Timeout: 3 * time.Minute} // Create a new HTTP client instance with a 3-minute timeout.

	// 1. Create a new request object to allow header customization (mimicking a browser)
	req, err := http.NewRequest("GET", finalURL, nil) // Create a new GET request object to allow manipulation of headers.
	if err != nil {                                   // Check for errors during the request object creation.
		log.Printf("Failed to create request for %s: %v", finalURL, err) // Log the error if request creation failed.
		return false                                                     // Return false.
	}

	// 2. Add comprehensive browser-like headers to increase download reliability
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36") // Set a realistic Chrome User-Agent string.
	req.Header.Set("Accept", "application/pdf,application/zip,application/octet-stream,*/*;q=0.8")                                                  // Set accepted content types, prioritizing PDF/ZIP/binary.
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")                                                                                          // Inform the server that the client supports these compression methods.
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")                                                                                             // Set the preferred language for the response.
	req.Header.Set("DNT", "1")                                                                                                                      // Set the Do Not Track header.
	req.Header.Set("Connection", "keep-alive")                                                                                                      // Indicate that the connection should remain open.

	// 3. Send the request using the client
	resp, err := client.Do(req) // Execute the configured request using the HTTP client.
	if err != nil {             // Check for network or connection errors during the request execution.
		log.Printf("Failed to download %s: %v", finalURL, err) // Log the error if the request failed.
		return false                                           // Return false.
	}
	// Make sure the response body gets closed when the function ends
	defer resp.Body.Close() // Ensure the raw response body stream is closed when the function exits.

	// Check that the HTTP status code is 200 OK
	if resp.StatusCode != http.StatusOK { // Check if the HTTP status code is not 200 OK.
		log.Printf("Download failed for %s: %s", finalURL, resp.Status) // Log the failure with the non-OK status.
		return false                                                    // Return false.
	}

	// Define the list of allowed content types
	allowedContentTypes := []string{ // Start defining a list of acceptable MIME types for the file.
		"application/pdf",             // PDF files.
		"application/zip",             // ZIP archives.
		"application/x-tar",           // TAR archives.
		"application/gzip",            // GZ files.
		"application/x-7z-compressed", // 7z archives.
		"application/vnd.rar",         // RAR archives.
		"application/octet-stream",    // General binary data (common fallback for archives).
	} // End of allowed content types list.

	// Get the Content-Type from the HTTP response headers
	contentType := resp.Header.Get("Content-Type") // Retrieve the Content-Type header value.

	// Flag to check if the content type is allowed
	isAllowed := false // Initialize a flag to track if the content type is valid.

	// Loop through allowed types and check for a match
	for _, allowed := range allowedContentTypes { // Iterate over the list of allowed types.
		if strings.Contains(contentType, allowed) { // Check if the response content type contains any allowed type (case-insensitive check).
			isAllowed = true // Set the flag to true if a match is found.
			break            // Stop checking once a match is found.
		}
	}

	// If the content type is not in the allowed list, skip download
	if !isAllowed { // Check if the content type validation failed.
		log.Printf("Invalid content type for %s %s (not allowed)", finalURL, contentType) // Log an error for an unexpected content type.
		return false                                                                      // Return false.
	}

	// --- FIX: Add decompression support for gzip/deflate if needed ---
	var reader io.ReadCloser                     // Declare an io.ReadCloser interface for the reading source (body or decompressor).
	switch resp.Header.Get("Content-Encoding") { // Check the Content-Encoding header to see if the body is compressed.
	case "gzip": // Case for Gzip compression.
		reader, err = gzip.NewReader(resp.Body) // Create a new Gzip reader that wraps the response body.
		if err != nil {                         // Check for errors when creating the Gzip reader.
			log.Printf("Failed to create gzip reader for %s %v", finalURL, err) // Log the error.
			return false                                                        // Return false.
		}
		defer reader.Close() // Ensure the Gzip reader is closed after reading.
	case "deflate": // Case for Deflate compression.
		reader = flate.NewReader(resp.Body) // Create a new Deflate reader that wraps the response body.
		defer reader.Close()                // Ensure the Deflate reader is closed after reading.
	default: // Default case (no compression or unsupported compression).
		reader = resp.Body // Use the raw response body as the reader source.
	}

	// Create a buffer to temporarily store the file in memory
	var buf bytes.Buffer // Declare a bytes.Buffer to temporarily hold the decompressed file data.

	// Read the entire response body (using the potentially decompressed reader) into the buffer
	written, err := io.Copy(&buf, reader) // Copy the data from the reader (decompressed or raw) into the buffer.
	if err != nil {                       // Check for errors during the read/copy process.
		log.Printf("Failed to read file data from %s %v", finalURL, err) // Log the error if reading the data failed.
		return false                                                     // Return false.
	}
	// --- END of Decompression and Read Fix ---

	// If no data was downloaded, skip file creation
	if written == 0 { // Check if zero bytes were copied into the buffer.
		log.Printf("Downloaded 0 bytes for %s; not creating file", finalURL) // Log a message if the downloaded size is zero.
		return false                                                         // Return false.
	}

	// Create a file on disk with the constructed file path
	out, err := os.Create(filePath) // Create the output file on the disk.
	if err != nil {                 // Check for errors during file creation.
		log.Printf("Failed to create file for %s %v", finalURL, err) // Log the error if file creation failed.
		return false                                                 // Return false.
	}
	// Ensure the file is properly closed at the end
	defer out.Close() // Ensure the output file handle is closed.

	// Write the contents from the buffer to the file on disk
	if _, err := buf.WriteTo(out); err != nil { // Write the content from the in-memory buffer to the file on disk.
		log.Printf("Failed to write file to disk for %s %v", finalURL, err) // Log the error if writing to the file failed.
		return false                                                        // Return false.
	}

	// Log the successful download
	log.Printf("Successfully downloaded %d bytes: %s → %s", written, finalURL, filePath) // Log a success message with the size and paths.
	return true                                                                          // Return true indicating success.
} // End of the downloadArchiveOrPDF function.
// Checks if the directory exists
// If it exists, return true.
// If it doesn't, return false.
func directoryExists(path string) bool { // Define a function to check if a directory exists, taking a string path and returning a boolean
	directory, err := os.Stat(path) // Get file information (os.FileInfo) for the path and any error
	if err != nil {                 // Check if an error occurred (e.g., file/directory doesn't exist)
		return false // If there was an error, return false
	} // End of the if block
	return directory.IsDir() // If no error, check if the retrieved file information indicates a directory and return the result
} // End of the directoryExists function

// The function takes two parameters: path and permission.
// We use os.Mkdir() to create the directory.
// If there is an error, we use log.Println() to log the error and then exit the program.
func createDirectory(path string, permission os.FileMode) { // Define a function to create a directory, taking a path string and file permissions (os.FileMode)
	err := os.Mkdir(path, permission) // Attempt to create the directory with the specified path and permissions
	if err != nil {                   // Check if an error occurred during directory creation
		log.Println(err) // Log the error to the standard logger
	} // End of the if block
} // End of the createDirectory function

// Checks whether a URL string is syntactically valid
func isUrlValid(uri string) bool { // Define a function to check URL validity, taking a string URI and returning a boolean
	_, err := url.ParseRequestURI(uri) // Attempt to parse the URI string as a URL; ignore the resulting URL structure and capture the error
	return err == nil                  // Return true if the error is nil (meaning parsing was successful), false otherwise
} // End of the isUrlValid function

// Remove all the duplicates from a slice and return the slice.
func removeDuplicatesFromSlice(slice []string) []string { // Define a function to remove duplicates from a string slice, taking a slice and returning a new slice
	check := make(map[string]bool)  // Create an empty map to keep track of seen elements (key=element, value=boolean existence)
	var newReturnSlice []string     // Declare an empty string slice to store unique elements
	for _, content := range slice { // Iterate over each element in the input slice
		if !check[content] { // Check if the current element (content) is NOT present in the map (i.e., it's a new element)
			check[content] = true                            // Mark the current element as seen in the map
			newReturnSlice = append(newReturnSlice, content) // Append the unique element to the new return slice
		} // End of the if block
	} // End of the loop
	return newReturnSlice // Return the slice containing only unique elements
} // End of the removeDuplicatesFromSlice function

// extractZIPUrls takes an input string and returns all ZIP URLs found within href attributes
func extractZIPUrls(input string) []string { // Define a function to extract ZIP URLs, taking an input string and returning a string slice
	// Regular expression to match href="...zip"
	re := regexp.MustCompile(`href="([^"]+\.zip)"`) // Compile a regular expression to find `href="..."` where the content ends in `.zip` and capture the content within quotes
	matches := re.FindAllStringSubmatch(input, -1)  // Find all non-overlapping matches and return a slice of slices of strings, where the inner slice contains the full match and submatches (the captured ZIP URL)

	var zipUrls []string            // Declare an empty string slice to store the extracted ZIP URLs
	for _, match := range matches { // Iterate over each match found
		if len(match) > 1 { // Check if the match slice has more than one element (meaning the captured group exists)
			zipUrls = append(zipUrls, match[1]) // Append the first captured group (the ZIP URL itself) to the zipUrls slice
		} // End of the if block
	} // End of the loop
	return zipUrls // Return the slice of extracted ZIP URLs
} // End of the extractZIPUrls function

// Append some string to a slice and than return the slice.
func appendToSlice(slice []string, content string) []string { // Define a function to append a string to a slice, taking a slice and a string, and returning the modified slice
	// Append the content to the slice
	slice = append(slice, content) // Append the given string 'content' to the end of the input 'slice'
	// Return the slice
	return slice // Return the resulting slice
} // End of the appendToSlice function

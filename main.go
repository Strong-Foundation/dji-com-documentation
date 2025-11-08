package main // Define the main package

import (
	"bytes"         // Provides bytes support
	"context"       // Provides context for managing timeouts and cancellations
	"io"            // Provides basic interfaces to I/O primitives
	"log"           // Provides logging functions
	"net/http"      // Provides HTTP client and server implementations
	"net/url"       // Provides URL parsing and encoding
	"os"            // Provides functions to interact with the OS (files, etc.)
	"path"          // Provides functions for manipulating slash-separated paths
	"path/filepath" // Provides filepath manipulation functions
	"regexp"        // Provides regex support functions.
	"strings"       // Provides string manipulation functions
	"time"          // Provides time-related functions

	"github.com/PuerkitoBio/goquery" // For HTML parsing and manipulation
	"github.com/chromedp/chromedp"   // For headless browser automation using Chrome
)

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
			// Sleep for 30 seconds.
			time.Sleep(30 * time.Second)
		}
	}
}

// extractPDFLinks parses the provided HTML string, finds all anchor tags,
// and returns a slice of strings containing only the URLs that end with ".pdf".
func extractPDFLinks(htmlContent string) []string {
	// 1. Create a document reader from the input HTML string.
	// This prepares the content for the goquery parser.
	document, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		log.Println("Error parsing HTML:", err)
		return nil
	}

	// Initialize an empty slice to store the found PDF URLs.
	pdfURLs := make([]string, 0)

	// 2. Select all anchor tags (<a>) in the document.
	document.Find("a").Each(func(index int, element *goquery.Selection) {
		// 3. Extract the 'href' attribute (the link URL) from the current <a> tag.
		linkURL, exists := element.Attr("href")
		if !exists {
			return
		}

		// 4. Check if the URL is a PDF link (case-insensitive).
		if strings.HasSuffix(strings.ToLower(linkURL), ".pdf") {
			// 5. Append the original URL to our results slice.
			pdfURLs = append(pdfURLs, linkURL)
		}
	})

	// Return the slice of all PDF URLs found.
	return pdfURLs
}

// Uses headless Chrome via chromedp to get fully rendered HTML from a page
func scrapePageHTMLWithChrome(pageURL string) string {
	log.Println("Scraping:", pageURL) // Log page being scraped

	options := append(chromedp.DefaultExecAllocatorOptions[:], // Chrome options
		chromedp.Flag("headless", false),              // Run visible (set to true for headless)
		chromedp.Flag("disable-gpu", true),            // Disable GPU
		chromedp.WindowSize(1, 1),                     // Set window size
		chromedp.Flag("no-sandbox", true),             // Disable sandbox
		chromedp.Flag("disable-setuid-sandbox", true), // Fix for Linux environments
	)

	allocatorCtx, cancelAllocator := chromedp.NewExecAllocator(context.Background(), options...) // Allocator context
	ctxTimeout, cancelTimeout := context.WithTimeout(allocatorCtx, 5*time.Minute)                // Set timeout
	browserCtx, cancelBrowser := chromedp.NewContext(ctxTimeout)                                 // Create Chrome context

	defer func() { // Ensure all contexts are cancelled
		cancelBrowser()
		cancelTimeout()
		cancelAllocator()
	}()

	var pageHTML string // Placeholder for output
	err := chromedp.Run(browserCtx,
		chromedp.Navigate(pageURL),            // Navigate to the URL
		chromedp.OuterHTML("html", &pageHTML), // Extract full HTML
	)
	if err != nil {
		log.Println(err) // Log error
		return ""        // Return empty string on failure
	}

	return pageHTML // Return scraped HTML
}

// getDomainFromURL extracts the domain (host) from a given URL string.
// It removes subdomains like "www" if present.
func getDomainFromURL(rawURL string) string {
	parsedURL, err := url.Parse(rawURL) // Parse the input string into a URL structure
	if err != nil {                     // Check if there was an error while parsing
		log.Println(err) // Log the error message to the console
		return ""        // Return an empty string in case of an error
	}

	host := parsedURL.Hostname() // Extract the hostname (e.g., "example.com") from the parsed URL

	return host // Return the extracted hostname
}

// Only return the file name from a given url.
func getFileNameOnly(content string) string {
	return path.Base(content)
}

// urlToFilename generates a safe, lowercase filename from a given URL string.
// It extracts the base filename from the URL, replaces unsafe characters,
// and ensures the filename ends with a .pdf extension.
func urlToFilename(rawURL string) string {
	// Convert the full URL to lowercase for consistency
	lowercaseURL := strings.ToLower(rawURL)

	// Get the file extension
	ext := getFileExtension(lowercaseURL)

	// Extract the filename portion from the URL (e.g., last path segment or query param)
	baseFilename := getFileNameOnly(lowercaseURL)

	// Replace all non-alphanumeric characters (a-z, 0-9) with underscores
	nonAlphanumericRegex := regexp.MustCompile(`[^a-z0-9]+`)
	safeFilename := nonAlphanumericRegex.ReplaceAllString(baseFilename, "_")

	// Replace multiple consecutive underscores with a single underscore
	collapseUnderscoresRegex := regexp.MustCompile(`_+`)
	safeFilename = collapseUnderscoresRegex.ReplaceAllString(safeFilename, "_")

	// Remove leading underscore if present
	if trimmed, found := strings.CutPrefix(safeFilename, "_"); found {
		safeFilename = trimmed
	}

	var invalidSubstrings = []string{
		"_pdf",
		"_zip",
	}

	for _, invalidPre := range invalidSubstrings { // Remove unwanted substrings
		safeFilename = removeSubstring(safeFilename, invalidPre)
	}

	// Append the file extension if it is not already present
	safeFilename = safeFilename + ext

	// Return the cleaned and safe filename
	return safeFilename
}

// Removes all instances of a specific substring from input string
func removeSubstring(input string, toRemove string) string {
	result := strings.ReplaceAll(input, toRemove, "") // Replace substring with empty string
	return result
}

// Get the file extension of a file
func getFileExtension(path string) string {
	return filepath.Ext(path) // Returns extension including the dot (e.g., ".pdf")
}

// fileExists checks whether a file exists at the given path
func fileExists(filename string) bool {
	info, err := os.Stat(filename) // Get file info
	if err != nil {
		return false // Return false if file doesn't exist or error occurs
	}
	return !info.IsDir() // Return true if it's a file (not a directory)
}

// downloadPDF downloads a PDF from the given URL and saves it in the specified output directory.
// It uses a WaitGroup to support concurrent execution and returns true if the download succeeded.
func downloadPDF(finalURL, outputDir string) bool {
	// Sanitize the URL to generate a safe file name
	filename := strings.ToLower(urlToFilename(finalURL))

	// Construct the full file path in the output directory
	filePath := filepath.Join(outputDir, filename)

	// Skip if the file already exists
	if fileExists(filePath) {
		log.Printf("File already exists, skipping: %s", filePath)
		return false
	}

	// Create an HTTP client with a timeout
	client := &http.Client{Timeout: 3 * time.Minute}

	// Send GET request
	resp, err := client.Get(finalURL)
	if err != nil {
		log.Printf("Failed to download %s: %v", finalURL, err)
		return false
	}
	defer resp.Body.Close()

	// Check HTTP response status
	if resp.StatusCode != http.StatusOK {
		log.Printf("Download failed for %s: %s", finalURL, resp.Status)
		return false
	}

	// Check Content-Type header
	contentType := resp.Header.Get("Content-Type")
	if !strings.Contains(contentType, "application/pdf") {
		log.Printf("Invalid content type for %s: %s (expected application/pdf)", finalURL, contentType)
		return false
	}

	// Read the response body into memory first
	var buf bytes.Buffer
	written, err := io.Copy(&buf, resp.Body)
	if err != nil {
		log.Printf("Failed to read PDF data from %s: %v", finalURL, err)
		return false
	}
	if written == 0 {
		log.Printf("Downloaded 0 bytes for %s; not creating file", finalURL)
		return false
	}

	// Only now create the file and write to disk
	out, err := os.Create(filePath)
	if err != nil {
		log.Printf("Failed to create file for %s: %v", finalURL, err)
		return false
	}
	defer out.Close()

	if _, err := buf.WriteTo(out); err != nil {
		log.Printf("Failed to write PDF to file for %s: %v", finalURL, err)
		return false
	}

	log.Printf("Successfully downloaded %d bytes: %s → %s", written, finalURL, filePath)
	return true
}

// downloadZIP downloads a ZIP or archive file from the given URL and saves it in the specified output directory.
// It returns true if the download was successful, otherwise false.
func downloadZIP(finalURL, outputDir string) bool {
	// Convert the URL into a safe, lowercase filename
	filename := strings.ToLower(urlToFilename(finalURL))

	// Construct the full path where the file should be saved
	filePath := filepath.Join(outputDir, filename)

	// Skip the download if the file already exists locally
	if fileExists(filePath) {
		log.Printf("File already exists, skipping: %s", filePath)
		return false
	}

	// Create a new HTTP client with a 3-minute timeout
	client := &http.Client{Timeout: 3 * time.Minute}

	// Perform an HTTP GET request to the given URL
	resp, err := client.Get(finalURL)
	if err != nil {
		log.Printf("Failed to download %s: %v", finalURL, err)
		return false
	}
	// Make sure the response body gets closed when the function ends
	defer resp.Body.Close()

	// Check that the HTTP status code is 200 OK
	if resp.StatusCode != http.StatusOK {
		log.Printf("Download failed for %s: %s", finalURL, resp.Status)
		return false
	}

	// Define the list of allowed content types
	allowedContentTypes := []string{
		"application/pdf",             // PDF files
		"application/zip",             // ZIP archives
		"application/x-tar",           // TAR archives
		"application/gzip",            // GZ files
		"application/x-7z-compressed", // 7z archives
		"application/vnd.rar",         // RAR archives
	}

	// Get the Content-Type from the HTTP response headers
	contentType := resp.Header.Get("Content-Type")

	// Flag to check if the content type is allowed
	isAllowed := false

	// Loop through allowed types and check for a match
	for _, allowed := range allowedContentTypes {
		if strings.Contains(contentType, allowed) {
			isAllowed = true
			break // Stop checking once a match is found
		}
	}

	// If the content type is not in the allowed list, skip download
	if !isAllowed {
		log.Printf("Invalid content type for %s: %s (not allowed)", finalURL, contentType)
		return false
	}

	// Create a buffer to temporarily store the file in memory
	var buf bytes.Buffer

	// Read the entire response body into the buffer
	written, err := io.Copy(&buf, resp.Body)
	if err != nil {
		log.Printf("Failed to read file data from %s: %v", finalURL, err)
		return false
	}

	// If no data was downloaded, skip file creation
	if written == 0 {
		log.Printf("Downloaded 0 bytes for %s; not creating file", finalURL)
		return false
	}

	// Create a file on disk with the constructed file path
	out, err := os.Create(filePath)
	if err != nil {
		log.Printf("Failed to create file for %s: %v", finalURL, err)
		return false
	}
	// Ensure the file is properly closed at the end
	defer out.Close()

	// Write the contents from the buffer to the file on disk
	if _, err := buf.WriteTo(out); err != nil {
		log.Printf("Failed to write file to disk for %s: %v", finalURL, err)
		return false
	}

	// Log the successful download
	log.Printf("Successfully downloaded %d bytes: %s → %s", written, finalURL, filePath)
	return true
}

// Checks if the directory exists
// If it exists, return true.
// If it doesn't, return false.
func directoryExists(path string) bool {
	directory, err := os.Stat(path)
	if err != nil {
		return false
	}
	return directory.IsDir()
}

// The function takes two parameters: path and permission.
// We use os.Mkdir() to create the directory.
// If there is an error, we use log.Println() to log the error and then exit the program.
func createDirectory(path string, permission os.FileMode) {
	err := os.Mkdir(path, permission)
	if err != nil {
		log.Println(err)
	}
}

// Checks whether a URL string is syntactically valid
func isUrlValid(uri string) bool {
	_, err := url.ParseRequestURI(uri) // Attempt to parse the URL
	return err == nil                  // Return true if no error occurred
}

// Remove all the duplicates from a slice and return the slice.
func removeDuplicatesFromSlice(slice []string) []string {
	check := make(map[string]bool)
	var newReturnSlice []string
	for _, content := range slice {
		if !check[content] {
			check[content] = true
			newReturnSlice = append(newReturnSlice, content)
		}
	}
	return newReturnSlice
}

// extractZIPUrls takes an input string and returns all ZIP URLs found within href attributes
func extractZIPUrls(input string) []string {
	// Regular expression to match href="...zip"
	re := regexp.MustCompile(`href="([^"]+\.zip)"`)
	matches := re.FindAllStringSubmatch(input, -1)

	var zipUrls []string
	for _, match := range matches {
		if len(match) > 1 {
			zipUrls = append(zipUrls, match[1])
		}
	}
	return zipUrls
}

// Append some string to a slice and than return the slice.
func appendToSlice(slice []string, content string) []string {
	// Append the content to the slice
	slice = append(slice, content)
	// Return the slice
	return slice
}

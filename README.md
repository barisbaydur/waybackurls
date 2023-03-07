# waybackurls

Bring all the URLs that the Wayback machine knows for one or more domain names.

Usage Example

```
 ▶ waybackurls -hostFile <file> -outFile
 ▶ waybackurls -host <host> -outFile
 ▶ waybackurls -host <host>
```
Use -hostFile to specify a file with a list of domains to check.<br/>
Use -outFile to save the results to a file by domain name. If not used, the results will be printed to the terminal.

Help

```
  -host string
        This flag will specify a single domain to check.
  -hostFile string
        This flag will specify a file with a list of domains to check.
  -outFile
        This flag will save the results to a file by domain name. If not used, the results will be printed to the terminal.
```

# IPAddress
Library for handling IP addresses and subnets, both IPv4 and IPv6

[View Project Page](https://seancfoley.github.io/IPAddress/)

[View Javadoc](https://seancfoley.github.io/IPAddress/IPAddress/apidocs/)

[In the Maven Central Repository](https://repo1.maven.org/maven2/com/github/seancfoley/ipaddress/)
- group id: com.github.seancfoley
- artifact id: [ipaddress](https://search.maven.org/search?q=ipaddress)
- versions: [2.0.2](https://search.maven.org/artifact/com.github.seancfoley/ipaddress/2.0.2/jar), [3.0.0](https://search.maven.org/artifact/com.github.seancfoley/ipaddress/3.0.0/jar), [4.3.0](https://search.maven.org/artifact/com.github.seancfoley/ipaddress/4.3.0/jar), [5.0.2](https://search.maven.org/artifact/com.github.seancfoley/ipaddress/5.0.2/jar)

Developed as an Eclipse project, the project files are checked in so it can be easily be imported into an Eclipse workspace.

Version | Notes
------- | -------------
[v1.0.1](https://github.com/seancfoley/IPAddress/releases/tag/v1.0.1) | Requires Java 6 or higher
[v2.0.2](https://github.com/seancfoley/IPAddress/releases/tag/v2.0.2) | Requires Java 8 or higher
[v3.0.0](https://github.com/seancfoley/IPAddress/releases/tag/v3.0.0) | Requires Java 8 or higher, features MAC address support, EUI-48 and EUI-64 MAC integration with IPv6, new address framework, new IP string formats parsed and produced, and other additions
[v4.3.0](https://github.com/seancfoley/IPAddress/releases/tag/v4.3.0) | Requires Java 8 or higher.  Features new prefix length handling.  IPv4-network/IPv6-subnet-router-anycast/zero-host addresses are interpreted as the prefix block subnet, while other prefixed addresses are individual addresses. There exists the option to preserve the version 3 behaviour.  Version 4.2.0 has additional methods for managing prefix blocks.  Version 4.3 features improved parsing performance and a change to increment(long) behaviour for subnets.
**[Latest Version v5.0.2](https://github.com/seancfoley/IPAddress/releases/tag/v5.0.2)** | Requires Java 8 or higher.  Support for Java 9, 10, 11 JPMS modules - the code is compiled with Java 8 but provides a Java 9 compiled module-info.class file.  Compatible with Android using Android API level 24 or higher.  [Delete the module-info](https://github.com/seancfoley/IPAddress/issues/16) when using Android, it should be ignored when using Java 8.  Version 5 features the addition of IPAddress sequential range classes IP\*AddressSeqRange, the reorganization of classes and interfaces in inet.ipaddr.format package to standard, large, and string subpackages, enhanced address block splitting and merging functionality, and the improved parsing performance introduced with version 4.3.0.  Other enhancements listed on the [releases page](https://github.com/seancfoley/IPAddress/releases/tag/v5.0.0)

Planned future additions: ports to [**TypeScript**](https://www.typescriptlang.org/) / [**JavaScript**](https://www.npmjs.com/) and [**Go**](https://golang.org/).

## Add to Intellij Project

1. Click File from the toolbar
1. Project Structure (CTRL + SHIFT + ALT + S on Windows/Linux, ⌘ + ; on Mac OS X)
1. Select Modules at the left panel
1. Dependencies tab
1. '+' → JARs or directories
1. Library...
1. From Maven...
1. Enter "ipaddress" into search bar, search
1. After waiting for Intellij to locate ipaddress on maven, select the ipaddress version you prefer


## Getting Started

### Java

starting with address or subnet strings

    String ipv6Str = "::/64";
    String ipv4Str = "1.2.255.4/255.255.0.0";
    try {
        IPAddress ipv6Address = new IPAddressString(ipv6Str).toAddress();
        IPAddress ipv4Address = new IPAddressString(ipv4Str).toAddress();
        // use addresses
    } catch (AddressStringException e) {
        String msg = e.getMessage();//detailed message indicating improper format in address string
        // handle improperly formatted address string
    }
    
starting with host name strings

    String hostPortStr = "[a:b:c:d:e:f:a:b]:8080";
	String hostServiceStr = "a.b.com:service";
	String hostAddressStr = "1.2.3.4";
	String dnsStr = "a.b.com";
	try {
	    HostName host = new HostName(hostPortStr);
	    InetSocketAddress socketAddress = host.asInetSocketAddress();
	    // use socket address
	        
	    host = new HostName(hostServiceStr);
	    socketAddress = host.asInetSocketAddress(service -> service.equals("service") ? 100 : null);
	    // use socket address
	        
	    host = new HostName(hostAddressStr);
	    IPAddress address = host.asAddress(); // does not resolve
	    // use address
	        
	    host = new HostName(dnsStr);
	    address = host.toAddress(); // resolves if necessary
	    // use address
	        
	} catch (HostNameException | UnknownHostException e) {
	    String msg = e.getMessage();
	    // handle improperly formatted host name or address string
	}

### Kotlin

starting with address or subnet strings, using exceptions for invalid formats

    val ipv6Str = "a:b:c:d::a:b/64"
    try {
        val ipv6AddressStr = IPAddressString(ipv6Str)
        val ipv6Addr = ipv6AddressStr.toAddress()
        // use address
        println(ipv6Addr) // a:b:c:d::a:b/64
    } catch(e: AddressStringException) {
        // handle improperly formatted address string
        println(e.message)
    }
    
 starting with address or subnet strings, using nullable types and safe calls to handle invalid or unexpected formats
 
    val ipv6v4Str = "a:b:c:d:e:f:1.2.3.4/112"
    val ipv6v4AddressStr = IPAddressString(ipv6v4Str)
    val ipAddr: IPAddress? = ipv6v4AddressStr.address
    println(ipAddr) // a:b:c:d:e:f:102:304/112
    
    val ipv4Addr = ipAddr?.toIPv6()?.embeddedIPv4Address
    println(ipv4Addr) // 1.2.3.4/16

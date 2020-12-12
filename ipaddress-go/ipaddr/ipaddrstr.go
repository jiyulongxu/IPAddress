package ipaddr

// A string that is used to identify a network host.

type HostIdentifierString interface {

	//static final char SEGMENT_VALUE_DELIMITER = ',';

	// provides a normalized String representation for the host identified by this HostIdentifierString instance
	ToNormalizedString() string

	//Validate() HostIdentifierException

	GetAddress() *Address

	ToAddress() (*Address, error)
}

var (
	_ HostIdentifierString = &IPAddressString{}
	_ HostIdentifierString = &MACAddressString{}
)

var defaultMACAddrParameters *macAddressStringParameters = &macAddressStringParameters{}

// NewIPAddressString constructs an IPAddressString that will parse the given string according to the given parameters
func NewMACAddressString(str string, params MACAddressStringParameters) *MACAddressString {
	return &MACAddressString{str: str, params: convertMACParams(params)}
}

type MACAddressString struct { //TODO needs its own file
	str    string
	params *macAddressStringParameters // when nil, defaultParameters is used
}

func (addrStr *MACAddressString) getParams() *macAddressStringParameters {
	params := addrStr.params
	if params == nil {
		params = defaultMACAddrParameters
		addrStr.params = params
	}
	return params
}

func (addrStr *MACAddressString) ToNormalizedString() string {
	//TODO MACAddressString
	return ""
}

func (addrStr *MACAddressString) GetAddress() *Address {
	//TODO MACAddressString
	return nil
}

func (addrStr *MACAddressString) ToAddress() (*Address, error) {
	//TODO MACAddressString
	return nil, nil
}

// NewIPAddressString constructs an IPAddressString that will parse the given string according to the given parameters
func NewIPAddressString(str string, params IPAddressStringParameters) *IPAddressString {
	// TODO you could make the conversion lazy, only done when needed, but not so sure it's worth it, the conversion should be  fast
	// but I am tempted.  you would need to stop passing params around so much and get it when needed in the parsing code.
	// But consider that few use them, and even fewer would not use the builder,
	////and even for those, they could convert to the builder-based one on their own
	// and even with lazy conversion, you might end up converting all the time
	var p *ipAddressStringParameters
	if params == nil {
		p = defaultIPAddrParameters
	} else {
		p = getPrivateParams(params)
	}
	return &IPAddressString{str: str, params: p}
}

var defaultIPAddrParameters *ipAddressStringParameters = &ipAddressStringParameters{}

type IPAddressString struct {
	str             string
	params          *ipAddressStringParameters // when nil, default parameters is used, never access this field directly
	addressProvider IPAddressProvider
}

func (ipAddrStr *IPAddressString) getParams() *ipAddressStringParameters {
	params := ipAddrStr.params
	if params == nil {
		params = defaultIPAddrParameters
		ipAddrStr.params = params
	}
	return params
}

func (ipAddrStr *IPAddressString) GetValidationOptions() IPAddressStringParameters {
	return ipAddrStr.getParams()
}

func (addrStr *IPAddressString) ToNormalizedString() string {
	//TODO IPAddressString
	return ""
}

//TODO we do want the three validate functions, they allow validation without address object creation
//func (addrStr *IPAddressString) Validate() HostIdentifierException {
//	return nil
//}

func (addrStr *IPAddressString) GetAddress() *IPAddress {
	if addrStr.addressProvider == nil || !addrStr.addressProvider.isInvalid() { // Avoid the exception the second time with this check
		addr, _ := addrStr.ToAddress() /* note the exception is cached, it is not lost forever */
		return addr
	}
	return nil
}

//
// error can be AddressStringException or IncompatibleAddressException
func (addrStr *IPAddressString) ToAddress() (*IPAddress, error) {
	//addrStr.validate() //call validate so that we throw consistently, cover type == INVALID, and ensure the addressProvider exists
	return addrStr.addressProvider.getProviderAddress()
}

///**
//	 * Validates that this string is a valid IPv4 address, and if not, throws an exception with a descriptive message indicating why it is not.
//	 * @throws AddressStringException
//	 */
//	public void validateIPv4() throws AddressStringException {
//		validate(IPVersion.IPV4);
//		checkIPv4Exception();
//	}
//
//	/**
//	 * Validates that this string is a valid IPv6 address, and if not, throws an exception with a descriptive message indicating why it is not.
//	 * @throws AddressStringException
//	 */
//	public void validateIPv6() throws AddressStringException {
//		validate(IPVersion.IPV6);
//		checkIPv6Exception();
//	}
//
//	/**
//	 * Validates that this string is a valid address, and if not, throws an exception with a descriptive message indicating why it is not.
//	 * @throws AddressStringException
//	 */
//	@Override
//	public void validate() throws AddressStringException {
//		validate(null);
//	}
//
//	private void checkIPv4Exception() throws AddressStringException {
//		IPVersion version = addressProvider.getProviderIPVersion();
//		if(version != null && version.isIPv6()) {
//			throw new AddressStringException("ipaddress.error.address.is.ipv6");
//		} else if(validateException != null) {
//			throw validateException;
//		}
//	}
//
//	private void checkIPv6Exception() throws AddressStringException {
//		IPVersion version = addressProvider.getProviderIPVersion();
//		if(version != null && version.isIPv4()) {
//			throw new AddressStringException("ipaddress.error.address.is.ipv4");
//		} else if(validateException != null) {
//			throw validateException;
//		}
//	}
//
//	private boolean isValidated(IPVersion version) throws AddressStringException {
//		if(!addressProvider.isUninitialized()) {
//			if(version == null) {
//				if(validateException != null) {
//					throw validateException; // the two exceptions are the same, so we can choose either one
//				}
//			} else if(version.isIPv4()) {
//				checkIPv4Exception();
//			} else if(version.isIPv6()) {
//				checkIPv6Exception();
//			}
//			return true;
//		}
//		return false;
//	}
//
//	protected HostIdentifierStringValidator getValidator() {
//		return Validator.VALIDATOR;
//	}
//
//	private void validate(IPVersion version) throws AddressStringException {
//		if(isValidated(version)) {
//			return;
//		}
//		synchronized(this) {
//			if(isValidated(version)) {
//				return;
//			}
//			//we know nothing about this address.  See what it is.
//			try {
//				addressProvider = getValidator().validateAddress(this);
//			} catch(AddressStringException e) {
//				validateException = e;
//				addressProvider = IPAddressProvider.INVALID_PROVIDER;
//				throw e;
//			}
//		}
//	}

func getPrivateParams(orig IPAddressStringParameters) *ipAddressStringParameters {
	if p, ok := orig.(*ipAddressStringParameters); ok {
		return p
	}
	return ToIPAddressStringParamsBuilder(orig).ToParams().(*ipAddressStringParameters)
}

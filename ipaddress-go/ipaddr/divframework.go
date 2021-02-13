package ipaddr

type AddressStringDivision interface {
}

type IPAddressStringDivision interface {
	AddressStringDivision
}

// AddressGenericDivision serves as common interface to all divisions, including large divisions (> 64 bits)
type AddressGenericDivision interface { //TODO rename GenericDivisionType or GenericDivision
	AddressItem
	AddressStringDivision

	getAddrType() addrType

	Equals(AddressGenericDivision) bool

	// to add: getDefaultTextualRadix()? getDigitCount(int radix)? getMaxDigitCount()? getWildcardString()?
}

// Represents any standard address division, all of which can be converted to/from AddressDivision
type AddressStandardDivision interface { //TODO rename StandardDivisionType
	AddressGenericDivision

	// GetDivisionValue returns the lower division value
	GetDivisionValue() DivInt

	// GetUpperDivisionValue returns the upper division value
	GetUpperDivisionValue() DivInt

	ToAddressDivision() *AddressDivision
}

var (
	_ AddressStandardDivision = &AddressDivision{}
	_ AddressStandardDivision = &AddressSegment{}
	_ AddressStandardDivision = &IPAddressSegment{}
	_ AddressStandardDivision = &IPv4AddressSegment{}
	_ AddressStandardDivision = &IPv6AddressSegment{}
	_ AddressStandardDivision = &MACAddressSegment{}
)

// euqivalent to AddressSegment on Java side, serves as common interface to all segments
type AddressStandardSegment interface { //TODO rename AddressSegmentType

	AddressStandardDivision

	Contains(AddressStandardSegment) bool

	// GetSegmentValue returns the lower segment value as a SegInt, the same value as the DivInt value returned by GetDivisionValue()
	GetSegmentValue() SegInt

	// GetUpperSegmentValue returns the upper segment value as a SegInt, the same value as the DivInt value returned by GetUpperDivisionValue()
	GetUpperSegmentValue() SegInt

	//ToAddressDivision() *AddressDivision
	ToAddressSegment() *AddressSegment
}

var (
	_ AddressStandardSegment = &AddressSegment{}
	_ AddressStandardSegment = &IPv6AddressSegment{}
	_ AddressStandardSegment = &MACAddressSegment{}
	_ AddressStandardSegment = &IPv4AddressSegment{}
)
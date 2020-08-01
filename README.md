# ewm-product
UUID, or GUID, is a 128 bit number used as an ID number.  The same number can be represented in different forms.  This package provides conversion between the hexadecimal form and a non-standard base64 form. 

Encode() converts hexadecimal to the non-standard base64.

Decode() converts the the non-standard base64 to hexadecimal.

** The non-standard-ness is the different character set used in the base64 index table.  Otherwise it is a copy from the standard library. See code for details.

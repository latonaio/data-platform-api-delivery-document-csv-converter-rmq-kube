package dpfm_api_output_formatter

import (
	"encoding/json"
	"os"

	"github.com/gocarina/gocsv"
	"golang.org/x/xerrors"
)

func ConvertToConcatMessage(filePath string) (*[]DataConcatenation, error) {
	data := make([]DataConcatenation, 0)
	f, err := os.OpenFile(filePath, os.O_RDONLY, 0)
	if err != nil {
		return nil, xerrors.Errorf("file open error: %w", err)
	}
	// defer f.Close()

	allHeaders := make([]DeliveryDocumentHeader, 0)
	allItems := make([]DeliveryDocumentItem, 0)
	allAddresses := make([]DeliveryDocumentAddress, 0)
	allPartners := make([]DeliveryDocumentPartner, 0)

	err = gocsv.UnmarshalFile(f, &allHeaders)
	if err != nil {
		return nil, xerrors.Errorf("read order header error: %w", err)
	}
	f.Close()
	f, _ = os.OpenFile(filePath, os.O_RDONLY, 0)
	err = gocsv.UnmarshalFile(f, &allItems)
	if err != nil {
		return nil, xerrors.Errorf("read order items error: %w", err)
	}
	f.Seek(0, 0)
	err = gocsv.UnmarshalFile(f, &allAddresses)
	if err != nil {
		return nil, xerrors.Errorf("read order addresses error: %w", err)
	}
	f.Seek(0, 0)
	err = gocsv.UnmarshalFile(f, &allPartners)
	if err != nil {
		return nil, xerrors.Errorf("read order partners error: %w", err)
	}

	allHeaders = toNull(allHeaders)
	allItems = toNull(allItems)
	allAddresses = toNull(allAddresses)
	allPartners = toNull(allPartners)
	allHeaders = getUniqueHeaders(&allHeaders)

	for headerIdx, v := range allHeaders {
		dID := v.DeliveryDocument
		items := getUniqueItem(&allItems, dID)
		if len(items) == 0 {
			return nil, xerrors.Errorf("items length is 0")
		}
		itemIDs := getItemIDs(items, dID)
		addresses := getUniqueAddresses(&allAddresses, dID, itemIDs)
		partners := getUniquePartners(&allPartners, dID)
		orderAll := DataConcatenation{
			Header:  allHeaders[headerIdx],
			Item:    items,
			Address: addresses,
			Partner: partners,
		}
		data = append(data, orderAll)
	}

	return &data, nil
}

func getOrderIDs(headers []DeliveryDocumentHeader) []int {
	idCheck := make(map[int]struct{})
	for _, v := range headers {
		idCheck[v.DeliveryDocument] = struct{}{}
	}
	ids := make([]int, 0)
	for k := range idCheck {
		ids = append(ids, k)
	}
	return ids
}

func getUniqueHeaders(headers *[]DeliveryDocumentHeader) []DeliveryDocumentHeader {
	idCheck := make(map[int]struct{})
	unique := make([]DeliveryDocumentHeader, 0, len(*headers))
	for i, v := range *headers {
		_, ok := idCheck[v.DeliveryDocument]
		if ok {
			continue
		}
		idCheck[v.DeliveryDocument] = struct{}{}
		unique = append(unique, (*headers)[i])
	}
	return unique
}

func getUniqueItem(items *[]DeliveryDocumentItem, dID int) []DeliveryDocumentItem {
	itemIDs := make([]int, 0)
	unique := make([]DeliveryDocumentItem, 0, len(*items))
	for i, v := range *items {
		if v.DeliveryDocument != dID || isContain(itemIDs, v.DeliveryDocumentItem) {
			continue
		}
		itemIDs = append(itemIDs, v.DeliveryDocumentItem)
		unique = append(unique, (*items)[i])
	}
	return unique
}

func getUniqueAddresses(addresses *[]DeliveryDocumentAddress, dID int, itemID []int) []DeliveryDocumentAddress {
	addIDs := make([]int, 0)
	unique := make([]DeliveryDocumentAddress, 0, len(*addresses))
	for i, v := range *addresses {
		if v.DeliveryDocument != dID || isContain(addIDs, v.AddressID) {
			continue
		}
		addIDs = append(addIDs, v.AddressID)
		unique = append(unique, (*addresses)[i])
	}
	return unique
}

func getUniquePartners(partners *[]DeliveryDocumentPartner, dID int) []DeliveryDocumentPartner {
	bpIDs := getBusinessPartnerIDs(partners, dID)
	unique := make([]DeliveryDocumentPartner, 0, len(*partners))
	for _, bpID := range bpIDs {
		partnerFuncs := make([]string, 0)
		for i, v := range *partners {
			if v.DeliveryDocument != dID || v.BusinessPartner != bpID {
				continue
			}

			if isContain(partnerFuncs, v.PartnerFunction) {
				continue
			}
			partnerFuncs = append(partnerFuncs, v.PartnerFunction)
			unique = append(unique, (*partners)[i])
		}
	}
	return unique
}

func getItemIDs(csv []DeliveryDocumentItem, dID int) []int {
	idCheck := make(map[int]struct{})
	for _, v := range csv {
		if v.DeliveryDocument != dID {
			continue
		}
		idCheck[v.DeliveryDocument] = struct{}{}
	}
	ids := make([]int, 0)
	for k := range idCheck {
		ids = append(ids, k)
	}
	return ids
}

func getBusinessPartnerIDs(partners *[]DeliveryDocumentPartner, dID int) []int {
	idCheck := make(map[int]struct{})
	for _, v := range *partners {
		if v.DeliveryDocument != dID {
			continue
		}

		idCheck[v.BusinessPartner] = struct{}{}
	}
	ids := make([]int, 0)
	for k := range idCheck {
		ids = append(ids, k)
	}
	return ids
}

func isContain[T comparable](obj []T, targ T) bool {
	for _, v := range obj {
		if v == targ {
			return true
		}
	}
	return false
}

func toNull[T any](obj []T) []T {
	b, _ := json.Marshal(obj)
	j := make([]map[string]interface{}, 0)
	json.Unmarshal(b, &j)
	for i := range j {
		for k, v := range j[i] {
			switch typedV := v.(type) {
			case int:
				if isZero(typedV) {
					j[i][k] = nil
				}
			case int32:
				if isZero(typedV) {
					j[i][k] = nil
				}
			case int64:
				if isZero(typedV) {
					j[i][k] = nil
				}
			case float32:
				if isZero(typedV) {
					j[i][k] = nil
				}
			case float64:
				if isZero(typedV) {
					j[i][k] = nil
				}
			case string:
				if isZero(typedV) {
					j[i][k] = nil
				}
			}
		}
	}
	b, _ = json.Marshal(j)
	json.Unmarshal(b, &obj)
	return obj
}
func isZero[T comparable](obj T) bool {
	var zero T
	return obj == zero
}

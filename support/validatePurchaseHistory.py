PRODUCT_ID_LENGTH = 5
CUSTOMER_ID_LENGTH = 6
ZIP_LENGTH = 5

print("Validating purchase history file. Problems found: ")

purchaseFile = open("/Volumes/Data Disk/Process App Data/Test/purchaseHistory.txt")

lineCount = 0

for line in purchaseFile:

    lineCount += 1

    tokens = line.split()

    if len(tokens) != 2:

        print("  > Line", lineCount, "has", len(tokens), "entries:", tokens)
    
    else:

        if len(tokens[0]) != CUSTOMER_ID_LENGTH or len(tokens[1]) != PRODUCT_ID_LENGTH:

            print("  > Line", lineCount, "has a bad entry:", tokens)

purchaseFile.close()



print("Validating product id file. Problems found: ")

productFile = open("/Volumes/Data Disk/Process App Data/Test/productIds.txt")

lineCount = 0

for line in productFile:

    lineCount += 1

    tokens = line.split()

    if len(tokens) != 1:

        print("  > Line", lineCount, "has", len(tokens), "entries:", tokens)
    
    else:

        if len(tokens[0]) != PRODUCT_ID_LENGTH:

            print("  > Line", lineCount, "has a bad entry:", tokens)

productFile.close()


print("Validating customer/zip file. Problems found: ")

customerZipFile = open("/Volumes/Data Disk/Process App Data/Test/customerIdsWithZips.txt")

lineCount = 0

for line in customerZipFile:

    lineCount += 1

    tokens = line.split()

    if len(tokens) != 2:

        print("  > Line", lineCount, "has", len(tokens), "entries:", tokens)

    else:

        if len(tokens[0]) != CUSTOMER_ID_LENGTH or len(tokens[1]) != ZIP_LENGTH:

            print("  > Line", lineCount, "has a bad entry:", tokens)

customerZipFile.close()

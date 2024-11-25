This is an example system. There are places where you might ask "why would you do it this way". The
answer is that this is a learning tool, a step on a journey. 

It's a bit fragile and needs correct data and some attention to relationships between some constants.

The system consists of a set of process chains that solve the following problem: Given a set of product ids, 
customer ids with associated zip codes and a purchase history (pairs of customerId <--> productId indicating
the customer bought the product), determine which zip code received a delivery farthest from some specified
zip code.

In order to run, you need data files. The data files are large and need to be generated. The C++ program in
the folder "support", generateDataFiles.cpp (which must be compiled with BinaryTree.cpp - 
"g++ generateDataFiles.cpp BinaryTree.cpp"), will generate three files and place them in the folder "data":

- productIds.txt : a set of product ids
- customerIdsWIthZips.txt : pairs of customerId <--> zipCode (where the customer lives)
- purchaseHistory.txt : pairs of customerId <--> productId (what the customer bought)

There are two files you can use to generate zip codes. This is specified by the constant ZIPCODE_FILE.

- zipCodes.txt: This file has zip codes including Hawaii and Alaska, but not various territories and
military bases (tryed to eliminate the obvious farthest zip codes, though HI is still there).

- zipCodesContiguous.txt: This file just has zip codes for the 48 contiguous states.

You can validate the files are in the correct format by running validateDataFiles.py. It ensures the files
have a data format expected by the system. It does not ensure the files are coherent among themselves (e.g., 
that all customer ids in the purchase history file are customer ids from the customer id/zip file.)

When you run, you can hardcode a zip into the main() variable, homeZip, or you can specify one on the command
line

```> ./main 15068```
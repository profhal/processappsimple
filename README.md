This is a work in progress. It's a bit fragile and needs correct data and some attention to relationships
between some constants.

The system consists of a set of process chains that solve the following problem: Given a set of product ids, 
customer ids with associated zip codes and a purchase history (pairs of customerId <--> productId indicating
the customer bought the product), determine which zip code received a delivery farthest from some specified
zip code.

In order to run, you need data files. The data files are large and need to be generated. The C++ program in
the folder "support", generateDataFiles.cpp (which must be compiled with BinaryTree.cpp - 
"g++ generateDataFiles.cpp BinaryTree.cpp"), will generate three files:

- productIds.txt : a set of product ids
- customerIdsWIthZips.txt : pairs of customerId <--> zipCode (where the customer lives)
- purchaseHistory.txt : pairs of customerId <--> productId (what the customer bought)

You can control the size of the files generated via constants in the C++ file (should be clear). At first
commit, the configuation was 1,000,000 product ids and 1,000,000 customer ids with associated zip codes.
The purchase history file size depends on the run as each customer is configured to buy between 24 and 72
items (based on internet research of Amazon users). So, exepct this file to have around 48M entries.

The data files can live wherever. You can indicate where the app will find these files in main.go by
changing the constant

const DATA_FOLDER_PATH = "&lt;path to the folder with the data files&gt;"

From there you should be able to build and run.

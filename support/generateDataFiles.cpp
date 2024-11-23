#include <chrono>
#include <cmath>
#include <ctime>
#include <fstream>
#include <iostream>
#include <string>

#include "BinaryTree.h"

using namespace std;

const string CHARACTER_BANK = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789";

string *generateProductIds(int idCount, int idLength,
                           bool showProgress = false) {

  BinaryTree idTree;

  string *result = new string[idCount];
  int resultCount = 0;

  string candidate;

  bool stillTrying;

  for (int p = 0; p < idCount; ++p) {

    if (showProgress && p % (idCount / 10) == 0) {

      cout << ".";
      cout.flush();
    }

    stillTrying = true;

    while (stillTrying) {

      candidate = "";

      // Starts with a letter
      //
      candidate += CHARACTER_BANK[rand() % 26];

      for (int c = 1; c < idLength; ++c) {

        candidate += CHARACTER_BANK[rand() % CHARACTER_BANK.length()];
      }

      if (!idTree.has(candidate)) {

        stillTrying = false;
      }
    }

    idTree.insert(candidate);
  }

  return idTree.getAscendingArray();
}

string *generateCustomerIds(int idCount, int idLength,
                            bool showProgress = false) {

  // Product and customer ids look the same.
  //
  return generateProductIds(idCount, idLength, showProgress);
}

const string ZIPCODE_FILE = "zipCodes.txt";

string *loadZipCodes(int &totalZips) {

  const int ZIP_STEP = 10000;
  string *zips = new string[ZIP_STEP];

  ifstream zipCodeFile;

  zipCodeFile.open(ZIPCODE_FILE);

  string token;

  totalZips = 0;

  while (zipCodeFile >> token) {

    zipCodeFile >> zips[totalZips++];

    getline(zipCodeFile, token);

    if (totalZips % ZIP_STEP == 0) {

      string *temp = new string[totalZips + ZIP_STEP];

      for (int z = 0; z < totalZips; ++z) {
        temp[z] = zips[z];
      }

      delete[] zips;

      zips = temp;
    }
  }

  return zips;
}

int main() {

  const int PRODUCT_COUNT = 100;
  const int PRODUCT_ID_LENGTH = 5;

  const int CUSTOMER_COUNT = 100; // made smaller constants for testing
  const int CUSTOMER_ID_LENGTH = 6;

  const int PURCHASES_PER_CUSTOMER_MIN = 24;
  const int PURCHASES_PER_CUSTOMER_MAX = 72;

  const string FILEPATH_TO_DATA_FOLDER = "../db/init_data/";
  const string PRODUCT_ID_FILENAME = "productIds";
  const string CUSTOMER_ID_WITH_ZIP_FILENAME = "customerIdsWithZips";
  const string PURCHASE_HISTORY_FILENAME = "purchaseHistory";

  const string DATA_FILE_EXTESION = ".txt";

  ofstream productIdFile;
  ofstream customerIdWithZipFile;
  ofstream purchaseHistoryFile;

  chrono::system_clock::time_point start;
  chrono::system_clock::time_point stop;
  chrono::duration<double> elapsedTime;

  int purchaseCount;

  string *zipBank;
  int zipBankCount = 0;
  string *productIds;
  string *customerIds;

  srand(time(0));

  zipBank = loadZipCodes(zipBankCount);

  cout << "Generating product ids ";
  cout.flush();

  start = chrono::system_clock::now();

  productIds = generateProductIds(PRODUCT_COUNT, PRODUCT_ID_LENGTH, true);

  stop = chrono::system_clock::now();

  elapsedTime = chrono::duration_cast<chrono::seconds>(stop - start);

  cout << "Done: " << elapsedTime.count() << "s." << endl;

  cout << "Generating customer ids ";
  cout.flush();

  start = chrono::system_clock::now();

  customerIds = generateProductIds(CUSTOMER_COUNT, CUSTOMER_ID_LENGTH, true);

  stop = chrono::system_clock::now();

  elapsedTime = chrono::duration_cast<chrono::seconds>(stop - start);

  cout << "Done: " << elapsedTime.count() << "s." << endl;

  productIdFile.open(FILEPATH_TO_DATA_FOLDER + PRODUCT_ID_FILENAME +
                     DATA_FILE_EXTESION);

  cout << "Writing product ids... ";
  cout.flush();

  start = chrono::system_clock::now();

  for (int p = 0; p < PRODUCT_COUNT; ++p) {

    productIdFile << productIds[p];

    if (p < PRODUCT_COUNT - 1) {

      productIdFile << endl;
    }
  }

  stop = chrono::system_clock::now();

  elapsedTime = chrono::duration_cast<chrono::seconds>(stop - start);

  cout << "Done: " << elapsedTime.count() << "s." << endl;

  productIdFile.close();

  cout << "Generating customer id files ";
  cout.flush();

  customerIdWithZipFile.open(FILEPATH_TO_DATA_FOLDER +
                             CUSTOMER_ID_WITH_ZIP_FILENAME +
                             DATA_FILE_EXTESION);

  start = chrono::system_clock::now();

  for (int c = 0; c < CUSTOMER_COUNT; ++c) {

    if (c % (CUSTOMER_COUNT / 10) == 0) {

      cout << ".";
      cout.flush();
    }

    customerIdWithZipFile << customerIds[c] << ","
                          << zipBank[rand() % zipBankCount];

    if (c < CUSTOMER_COUNT - 1) {

      customerIdWithZipFile << endl;
    }
  }

  stop = chrono::system_clock::now();

  elapsedTime = chrono::duration_cast<chrono::seconds>(stop - start);

  cout << "Done: " << elapsedTime.count() << "s." << endl;

  customerIdWithZipFile.close();

  delete[] zipBank;

  cout << "Generating purchase history files ";
  cout.flush();

  purchaseHistoryFile.open(FILEPATH_TO_DATA_FOLDER + PURCHASE_HISTORY_FILENAME +
                           DATA_FILE_EXTESION);

  start = chrono::system_clock::now();

  for (int c = 0; c < CUSTOMER_COUNT; ++c) {

    if (c % (CUSTOMER_COUNT / 10) == 0) {

      cout << ".";
      cout.flush();
    }

    purchaseCount = (rand() % (PURCHASES_PER_CUSTOMER_MAX + 1 -
                               PURCHASES_PER_CUSTOMER_MIN) +
                     PURCHASES_PER_CUSTOMER_MIN);

    for (int p = 0; p < purchaseCount; ++p) {

      purchaseHistoryFile << customerIds[c] << ","
                          << productIds[rand() % PRODUCT_COUNT];

      if (p < purchaseCount - 1) {

        purchaseHistoryFile << endl;
      }
    }
  }

  stop = chrono::system_clock::now();

  elapsedTime = chrono::duration_cast<chrono::seconds>(stop - start);

  cout << "Done: " << elapsedTime.count() << "s." << endl;

  purchaseHistoryFile.close();

  delete[] productIds;
  delete[] customerIds;

  return 0;
}

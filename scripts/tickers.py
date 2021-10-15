from qtrade_client.api import QtradeAPI, APIException
from pprint import pprint

if __name__ == "__main__":
    api = QtradeAPI("https://api.qtrade.io", key="1:1111")

    for _, ticker in api.tickers.items():
        print(str(ticker["id_hr"]) + ": " + str(ticker["id"]))
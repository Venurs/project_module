from handlers.LoginHandler import LoginHandler, IndexHandler
from handlers.OrderingBalanceHandler import OrderingBalanceHandler, RefundInfoHandler, CompanyMasterHandler
from handlers.CollectionHandler import CollectionHandler, PayAccountHandler, AccountListHandler
from handlers.ExRateHandler import ExchangeRateHandler, CurrencyHandler
from handlers.PaymentsDetailsHandler import PaymentsDetailsHandler
from handlers.SupplierBalanceHandler import SupplierBalanceHandler
from handlers.SupplierPaymentHandler import SupplierPaymentHandler, SupplierList, CompanyBodyList
from handlers.WorkOrderHandler import WorkOrderHandler
from handlers.ProductHandler import ProductHandler
from handlers.CustomerHandler import CustomerHandler

urls = [
    (r'/cwapi/orderbalance/', OrderingBalanceHandler),
    (r'/cwapi/refundinfo/', RefundInfoHandler),
    (r'/cwapi/companymaster/', CompanyMasterHandler),
    (r'/cwapi/collectinfo/', CollectionHandler),
    (r'/cwapi/rate/', ExchangeRateHandler),
    (r'/cwapi/rate/currency/', CurrencyHandler),
    (r'/cwapi/order/work_summary/', WorkOrderHandler),
    (r'/cwapi/payaccount/', PayAccountHandler),
    (r'/cwapi/product/', ProductHandler),
    (r'/cwapi/customer/', CustomerHandler),
    (r'/cwapi/paydetails/', PaymentsDetailsHandler),
    (r'/cwapi/supplierbalance/', SupplierBalanceHandler),
    (r'/cwapi/supplierpayment/', SupplierPaymentHandler),
    (r'/cwapi/supplierlist/', SupplierList),
    (r'/cwapi/companybodylist/', CompanyBodyList),
    (r'/cwapi/accountlist/', AccountListHandler),
]





































# /**
#  *                    .::::.
#  *                  .::::::::.
#  *                 :::::::::::  FUCK YOU
#  *             ..:::::::::::'
#  *           '::::::::::::'
#  *             .::::::::::
#  *        '::::::::::::::..
#  *             ..::::::::::::.
#  *           ``::::::::::::::::
#  *            ::::``:::::::::'        .:::.
#  *           ::::'   ':::::'       .::::::::.
#  *         .::::'      ::::     .:::::::'::::.
#  *        .:::'       :::::  .:::::::::' ':::::.
#  *       .::'        :::::.:::::::::'      ':::::.
#  *      .::'         ::::::::::::::'         ``::::.
#  *  ...:::           ::::::::::::'              ``::.
#  * ```` ':.          ':::::::::'                  ::::..
#  *                    '.:::::'                    ':'````..
#  */





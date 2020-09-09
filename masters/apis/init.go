package apis

import (
	"database/sql"

	"github.com/gorilla/mux"
	orderControllers "github.com/inact25/PickMyFood-BackEnd/masters/apis/controllers/order"
	paymentControllers "github.com/inact25/PickMyFood-BackEnd/masters/apis/controllers/payment"
	productControllers "github.com/inact25/PickMyFood-BackEnd/masters/apis/controllers/product"
	productCategoryControllers "github.com/inact25/PickMyFood-BackEnd/masters/apis/controllers/productCategory"
	storeControllers "github.com/inact25/PickMyFood-BackEnd/masters/apis/controllers/store"
	storeCategoryControllers "github.com/inact25/PickMyFood-BackEnd/masters/apis/controllers/storeCategory"
	userControllers "github.com/inact25/PickMyFood-BackEnd/masters/apis/controllers/user"
	walletControllers "github.com/inact25/PickMyFood-BackEnd/masters/apis/controllers/wallet"
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/middlewares"
	orderRepositories "github.com/inact25/PickMyFood-BackEnd/masters/apis/repositories/order"
	paymentRepositories "github.com/inact25/PickMyFood-BackEnd/masters/apis/repositories/payment"
	productRepositories "github.com/inact25/PickMyFood-BackEnd/masters/apis/repositories/product"
	productCategoryRepositories "github.com/inact25/PickMyFood-BackEnd/masters/apis/repositories/productCategory"
	storerepositories "github.com/inact25/PickMyFood-BackEnd/masters/apis/repositories/store"
	storeCategoryRepositories "github.com/inact25/PickMyFood-BackEnd/masters/apis/repositories/storeCategory"
	userRepositories "github.com/inact25/PickMyFood-BackEnd/masters/apis/repositories/user"
	walletrepositories "github.com/inact25/PickMyFood-BackEnd/masters/apis/repositories/wallet"
	orderUsecases "github.com/inact25/PickMyFood-BackEnd/masters/apis/usecases/order"
	paymentUsecases "github.com/inact25/PickMyFood-BackEnd/masters/apis/usecases/payment"
	productUsecases "github.com/inact25/PickMyFood-BackEnd/masters/apis/usecases/product"
	productCategoryUsecases "github.com/inact25/PickMyFood-BackEnd/masters/apis/usecases/productCategory"
	storeusecases "github.com/inact25/PickMyFood-BackEnd/masters/apis/usecases/store"
	storeCategoryUsecases "github.com/inact25/PickMyFood-BackEnd/masters/apis/usecases/storeCatergory"
	userUsecases "github.com/inact25/PickMyFood-BackEnd/masters/apis/usecases/user"
	walletusecases "github.com/inact25/PickMyFood-BackEnd/masters/apis/usecases/wallet"
)

func Init(r *mux.Router, db *sql.DB) {
	// user
	userRepo := userRepositories.InitUserRepoImpl(db)
	userUseCases := userUsecases.InitUsersUseCase(userRepo)
	usersController := userControllers.UsersController(userUseCases)
	usersController.Authenticate(r)
	// wallet
	walletRepo := walletrepositories.InitWalletRepoImpl(db)
	walletUseCases := walletusecases.InitWalletUseCase(walletRepo)
	walletController := walletControllers.WalletController(walletUseCases)
	walletController.WalletApi(r)
	// store
	storeRepo := storerepositories.InitStoreRepoImpl(db)
	storeUseCase := storeusecases.InitStoreUseCase(storeRepo)
	storeController := storeControllers.StoreController(storeUseCase)
	storeController.StoreAPI(r)
	// storeCategory
	storeCategoryRepo := storeCategoryRepositories.InitStoreCategoryRepoImpl(db)
	storeCategoryUseCase := storeCategoryUsecases.InitStoreCategoryUseCase(storeCategoryRepo)
	storeCategoryController := storeCategoryControllers.StoreCategoryController(storeCategoryUseCase)
	storeCategoryController.StoreCategoryAPI(r)
	// productCategory
	productCategoryRepo := productCategoryRepositories.InitProductCategoryRepoImpl(db)
	productCategoryUseCase := productCategoryUsecases.InitProductCategoryUseCase(productCategoryRepo)
	productCategoryController := productCategoryControllers.ProductCategoryController(productCategoryUseCase)
	productCategoryController.ProductCategoryAPI(r)
	// product
	productRepo := productRepositories.InitProductRepoImpl(db)
	productUseCase := productUsecases.InitProductUseCaseImpl(productRepo)
	productController := productControllers.InitProductController(productUseCase)
	productController.ProductAPI(r)
	// order
	orderRepo := orderRepositories.InitOrderRepoImpl(db)
	orderUseCase := orderUsecases.InitOrderUseCaseImpl(orderRepo)
	orderController := orderControllers.InitOrderController(orderUseCase)
	orderController.OrderAPI(r)
	// payment
	paymentRepo := paymentRepositories.InitPaymentRepoImpl(db)
	paymentUseCase := paymentUsecases.InitPaymentUseCaseImpl(paymentRepo)
	paymentController := paymentControllers.InitPaymentController(paymentUseCase)
	paymentController.PaymentAPI(r)

	r.Use(middlewares.ActivityLogMiddleware)
}

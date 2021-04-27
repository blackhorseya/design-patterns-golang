package abstractfactory

// OrderMainDAO declare main order's record
type OrderMainDAO interface {
	SaveOrderMain() string
}

// OrderDetailDAO declare detail order's record
type OrderDetailDAO interface {
	SaveOrderDetail() string
}

// DAOFactory declare DAO abstract factory interface
type DAOFactory interface {
	CreateOrderMainDAO() OrderMainDAO
	CreateOrderDetailDAO() OrderDetailDAO
}

// RDBMainDAO declare order main dao of relationship database
type RDBMainDAO struct {
}

// SaveOrderMain serve caller to save main order to rdb
func (*RDBMainDAO) SaveOrderMain() string {
	return "rdb main save"
}

// RDBDetailDAO declare order detail dao of relationship database
type RDBDetailDAO struct {
}

// SaveOrderDetail serve caller to save detail order to rdb
func (*RDBDetailDAO) SaveOrderDetail() string {
	return "rdb detail save"
}

// RDBDAOFactory declare rdb abstract factory implement
type RDBDAOFactory struct {
}

// CreateOrderMainDAO serve caller to create an OrderMainDAO
func (*RDBDAOFactory) CreateOrderMainDAO() OrderMainDAO {
	return &RDBMainDAO{}
}

// CreateOrderDetailDAO serve caller to create an OrderDetailDAO
func (*RDBDAOFactory) CreateOrderDetailDAO() OrderDetailDAO {
	return &RDBDetailDAO{}
}

// XMLMainDAO declare order main dao of xml
type XMLMainDAO struct {
}

// SaveOrderMain serve caller to save main order to xml
func (*XMLMainDAO) SaveOrderMain() string {
	return "xml main save"
}

// XMLDetailDAO declare order detail dao of xml
type XMLDetailDAO struct {
}

// SaveOrderDetail serve caller to save detail order to xml
func (*XMLDetailDAO) SaveOrderDetail() string {
	return "xml detail save"
}

// XMLDAOFactory declare xml abstract factory implement
type XMLDAOFactory struct {
}

// CreateOrderMainDAO serve caller to create an OrderMainDAO
func (*XMLDAOFactory) CreateOrderMainDAO() OrderMainDAO {
	return &XMLMainDAO{}
}

// CreateOrderDetailDAO serve caller to create an OrderDetailDAO
func (*XMLDAOFactory) CreateOrderDetailDAO() OrderDetailDAO {
	return &XMLDetailDAO{}
}

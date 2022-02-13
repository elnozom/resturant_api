USE RMSS

IF OBJECT_ID('NozTrCart', 'U') IS NULL
BEGIN
CREATE TABLE NozTrCart (
    CartSerial int IDENTITY(1,1) NOT NULL,
    TableSerial int NOT NULL,
    Amount real DEFAULT 0,
    CustomerSerial int DEFAULT NULL,
    DeviceId VARCHAR(100) NOT NULL
)
END
IF OBJECT_ID('NozTrCartItems', 'U') IS NULL
BEGIN
CREATE TABLE NozTrCartItems (
    CartItemSerial int IDENTITY(1,1) NOT NULL,
    CartSerial int NOT NULL,
    ItemSerial int NOT NULL,
    MainModSerial int DEFAULT 0,
    IsMod int DEFAULT 0,
    AddItems TEXT DEFAULT '',
    Qnt int NOT NULL DEFAULT 1,
    Price real NOT NULL
);
END



GO
EXEC DropProcIfExist @Name = "CartCreate"
GO
CREATE PROC CartCreate (@CustomerSerial INT = 0 , @DeviceId VARCHAR(100) , @TableSerial INT)
AS 
BEGIN
    INSERT INTO NozTrCart (CustomerSerial , DeviceId , TableSerial) VALUES (@CustomerSerial , @DeviceId , @TableSerial)

    SELECT 1 AS Created
END

GO
EXEC DropProcIfExist @Name = "CartClose"
GO
CREATE PROC CartClose (@Serial int)
AS 
BEGIN
    DECLARE @Amount real 
    SELECT @Amount = SUM(Qnt * Price) FROM NozTrCartItems WHERE CartSerial = @Serial
   UPDATE NozTrCart SET Amount = @Amount WHERE CartSerial = @Serial
    SELECT 1 AS Closed
END

GO
EXEC DropProcIfExist @Name = "CartItemList"
GO
CREATE PROC CartItemList (@Serial INT)
AS 
BEGIN
   SELECT ci.CartItemSerial , ci.Qnt , ISNULL(Price , 0) Price ,
	 ci.ItemSerial , i.ItemName ,ci.IsMod,ISNULL(ci.MainModSerial , 0) MainModSerial
	  , ISNULL(ci.AddItems , '') AddItems
	   FROM NozTrCart c 
       JOIN NozTrCartItems ci ON c.CartSerial = ci.CartSerial  
       JOIN StkMs01 i ON ci.ItemSerial = i.Serial  
	   WHERE c.CartSerial = @Serial
	   order by ci.CartItemSerial  
END


GO
EXEC DropProcIfExist @Name = "CartItemCreate"
GO
CREATE PROC CartItemCreate (@CartSerial INT , @ItemSerial INT, @Price REAL )
AS 
BEGIN
    INSERT INTO NozTrCartItems (CartSerial,ItemSerial,Price) VALUES (@CartSerial , @ItemSerial ,@Price)
    SELECT 1 AS Created
END




GO
EXEC DropProcIfExist @Name = "CartItemDelete"
GO
CREATE PROC CartItemDelete (@Serial INT)
AS 
BEGIN
    DELETE FROM NozTrCartItems WHERE CartItemSerial = @Serial
    SELECT 1 AS Deleted

END

GO
EXEC DropProcIfExist @Name = "CartItemUpdae"
EXEC DropProcIfExist @Name = "CartItemUpdate"
GO
CREATE PROC CartItemUpdate (@Serial INT , @Qnt INT)
AS 
BEGIN
    UPDATE NozTrCartItems SET Qnt = @Qnt WHERE CartItemSerial = @Serial
    SELECT 1 AS Updated
END
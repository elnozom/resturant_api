USE RMSS



UPDATE StkMs01  SET ImagePath = CONCAT((SELECT g.GroupTypeID FROM GroupCode g WHERE g.GroupCode = StkMs01.GroupCode), '/', GroupCode, '/' , BarCode , '.jpg')
SELECT GroupTypeID  FROM GroupType gt  

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

    SELECT SCOPE_IDENTITY() AS "Serial"
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
CREATE PROC CartItemList (@Table INT , @DeviceId VARCHAR(100))
AS 
BEGIN
   SELECT c.CartSerial , ci.CartItemSerial , ci.Qnt , ISNULL(Price , 0) Price ,
	 ci.ItemSerial , i.ItemName ,ci.IsMod,ISNULL(ci.MainModSerial , 0) MainModSerial
	  , ISNULL(ci.AddItems , '') AddItems
	   FROM NozTrCart c 
       JOIN NozTrCartItems ci ON c.CartSerial = ci.CartSerial  
       JOIN StkMs01 i ON ci.ItemSerial = i.Serial  
	   WHERE c.TableSerial = @Table AND c.DeviceId = @DeviceId
	   order by ci.CartItemSerial  
END


GO
EXEC DropProcIfExist @Name = "CartItemCreate"
GO
CREATE PROC CartItemCreate (@CartSerial INT , @ItemSerial INT, @Price REAL )
AS 
BEGIN
    INSERT INTO NozTrCartItems (CartSerial,ItemSerial,Price) VALUES (@CartSerial , @ItemSerial ,@Price)
    SELECT SCOPE_IDENTITY() AS "Serial"
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
EXEC DropProcIfExist @Name = "CartItemUpdate"
GO
CREATE PROC CartItemUpdate (@Serial INT , @Qnt INT)
AS 
BEGIN
    UPDATE NozTrCartItems SET Qnt = @Qnt WHERE CartItemSerial = @Serial
    SELECT 1 AS Updated
END


IF OBJECT_ID('NozCartCalls', 'U') IS NULL
BEGIN
CREATE TABLE NozCartCalls (
    NozCartCallsSerial int IDENTITY(1,1) NOT NULL,
    CartSerial int NOT NULL,
    CreatedAt DATETIME DEFAULT GETDATE(),
    RespondedAt DATETIME,
    WaiterCode INT,
    CallType bit 
)
END


GO
EXEC DropProcIfExist @Name = "CartCallCreate"
GO
CREATE PROC CartCallCreate (@CallType BIT , @CartSerial INT)
AS 
BEGIN
    INSERT INTO NozCartCalls (CallType , CartSerial) VALUES (@CallType , @CartSerial)
    SELECT SCOPE_IDENTITY() AS "Serial"
END



GO
EXEC DropProcIfExist @Name = "CartCallRespond"
GO
CREATE PROC CartCallRespond (@Serial INT , @WaiterCode INT)
AS 
BEGIN
    UPDATE NozCartCalls SET RespondedAt = GETDATE() ,  WaiterCode = @WaiterCode WHERE NozCartCallsSerial = @Serial
    SELECT 1 AS upadted
END



-- GO
-- EXEC DropProcIfExist @Name = "CartCheckCalls"
-- GO
-- CREATE PROC CartCheckCalls (@EmpCode INT)
-- AS 
-- BEGIN
   
-- END

ALTER TABLE StkMs01
ADD ItemNameEn VARCHAR(100); 


ALTER TABLE GroupCode
ADD GroupNameEn VARCHAR(100); 


ALTER TABLE GroupType
ADD GroupTypeNameEn VARCHAR(100);



ALTER TABLE StkMs01
ADD ImagePath VARCHAR(100); 
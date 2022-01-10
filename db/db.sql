

--- schema
	--employee--
	-- EmployeeGetByCode(@EmpCode int) :{Employee} [EmpName, EmpPassword , EmpCode]
	--tables--
	-- GroupTablesList :{GroupTables} [ GroupTableNo , GroupTableName , TableCount ]
	-- TablesListByGroupNo(@GroupTableNo int) :{Tables} [TableNo , TableName , "pause" , ISNULL(OpenDateTime , '' )]
	--groups--
	-- GroupTypeList :{GroupType} [GroupTypeID , GroupTypeName]
	-- GroupCodeListByGroupTypeId(@GroupTypeID int) :{GroupCode}	[g.GroupCode , g.GroupName ]
	--items--
	-- StkMs01ListByMenuAndGroup(@GroupCode int , @MenuSerial int = 1) : { ItemMenuMap  im , StkMs01 i } [im.ItemSerial , im.ItemPrice , i.ItemCode , i.ItemName , i.WithModifier]


USE RMSS

GO
CREATE PROCEDURE EmployeeGetByCode (@EmpCode int)
AS
BEGIN
	SELECT       EmpName, EmpPassword , EmpCode
	FROM            Employee 
	WHERE EmpCode = @EmpCode
END



GO
CREATE PROCEDURE GroupTablesList
AS
BEGIN
	-- SET NOCOUNT ON added to prevent extra result sets from
	-- interfering with SELECT statements.
	SET NOCOUNT ON;

	SELECT GroupTableNo , GroupTableName , TableCount FROM  GroupTables
END




GO
CREATE PROCEDURE [dbo].[TablesListByGroupNo](@GroupTableNo int)
AS
BEGIN
	-- SET NOCOUNT ON added to prevent extra result sets from
	-- interfering with SELECT statements.
	SET NOCOUNT ON;
	DECLARE @TotalCash FLOAT
	SELECT  Tables.Serial ,  Tables.TableNo , TableName , "pause" , "State" , ISNULL(StkTr03.PrintTimes , 0) PrintTimes , ISNULL(convert(char(5), DocDate, 108) , '') DocDate , ISNULL(StkTr03.DocNo , '') DocNo , ISNULL(StkTr03.Serial , 0) HeadSerail ,  ISNULL(StkTr03.WaiterCode ,0) WaiterCode , ISNULL((SELECT SUM(Qnt * Price) FROM  StkTr04 WHERE HeadSerial = StkTr03.Serial) ,0) TotalCash  FROM  "Tables" LEFT JOIN StkTr03 ON StkTr03.TableSerial = Tables.Serial AND ISNULL(TotalCash ,0) = 0 WHERE Tables.GroupTableNo = @GroupTableNo
END

GO
-- list all main groups
-- GroupType table contains the main groups
CREATE PROCEDURE GroupTypeList
AS
BEGIN
	-- SET NOCOUNT ON added to prevent extra result sets from
	-- interfering with SELECT statements.
	SET NOCOUNT ON;

	SELECT  GroupTypeID , GroupTypeName FROM  "GroupType"
END


GO
-- list all main groups
-- GrouCode table contains the sub groups 
-- this table has relationship with GroupType['main groups table']
CREATE PROCEDURE GroupCodeListByGroupTypeId(@GroupTypeID int)
AS
BEGIN
	-- SET NOCOUNT ON added to prevent extra result sets from
	-- interfering with SELECT statements.
	SET NOCOUNT ON;

	SELECT  g.GroupCode , g.GroupName  FROM  "GroupCode" g WHERE g.GroupTypeID = @GroupTypeID
END



-- list all main groups
-- GrouCode table contains the sub groups 
-- this table has relationship with GroupType['main groups table']
GO
CREATE PROCEDURE StkMs01ListByMenuAndGroup(@GroupCode int , @MenuSerial int = 1)
AS
BEGIN
	-- SET NOCOUNT ON added to prevent extra result sets from
	-- interfering with SELECT statements.
	SET NOCOUNT ON;

	SELECT im.ItemSerial , im.ItemPrice , i.ItemCode , i.ItemName , i.WithModifier 
		FROM ItemMenuMap  im JOIN StkMs01 i ON im.ItemSerial = i.Serial AND i.GroupCode = @GroupCode
	 WHERE im.MnuSerial = @MenuSerial
END




GO
-- insert order in stktr03 table
CREATE  PROCEDURE  [dbo].[Stktr03Insert] (@TableNo int, @GroupTableNo int,@Imei VARCHAR(50),@OrderType int,@WaiterCode int)
 AS
	
	DECLARE @Captial CHAR
	DECLARE @CashTryNo INT
	DECLARE @DocNo VARCHAR(20)
	DECLARE @BonNo INT
	DECLARE @OrderNo INT	
	DECLARE @ComName VARCHAR(20)
	DECLARE @CashSerial INT
	DECLARE @CashNo INT
	DECLARE @SessionNo INT
	DECLARE @TrSerial INT
	DECLARE @StoreCode INT
	DECLARE @AccSerial INT
	DECLARE @CasherCode INT
	DECLARE @GroupTableNo int
	DECLARE @TableNo int
	declare @CaherCashTrySerial int 
	declare @EmpCode   int 
	SELECT @GroupTableNo = GroupTableNo , @TableNo = TableNo FROM dbo.Tables WHERE Serial = @TableSerial
	SELECT 
	@ComName = ComName ,
	@CashSerial = CashTry.Serial ,
	@CashNo = CashTryNo ,
	@SessionNo = SessionNo ,
	@TrSerial = TransSerial ,
	@StoreCode = POSOptions.StoreCode ,
	@AccSerial = POSOptions.AccountSerial ,
	@Captial =  Capital ,
	@CasherCode = ComUse.UserId,
	@CashTryNo = CashTryNo ,
	@CaherCashTrySerial = CasherCashTrySerial,
	@EmpCode = EmpCode
	FROM ComUse 
	 	JOIN CashTry ON ComUse.ComName = ComputerName
		,POSOptions
		
   WHERE ComUse.Imei = @Imei and CashTry.CloseDate is null and Paused = 0


	SELECT @OrderNo = ISNULL( MAX(OrderNo),0) + 1 FROM StkTr03
	SELECT  @BonNo =  ISNULL (MAX(BonNo),0) + 1 FROM StkTr03 WHERE CashTrySerial = @CashSerial




	SET @DocNo = CONCAT(@BonNo , '-' ,@Captial,@CashTryNo,'-',@SessionNo)



	INSERT INTO StkTr03 (
		DocNo ,
		TransSerial,
		CashTrySerial ,
		StoreCode ,
		TableNo,
		GroupTableNo ,
		AccountSerial ,
		OrderNo ,
		OrderType , 
		ComputerName,
		CasherCode ,
		BonNo ,
		WaiterCode , 
		TableSerial,
		CasherCashTrySerial
	) VALUES (
		@DocNo ,
		@TrSerial,
		@CashSerial,
		@StoreCode,
		@TableNo,
		@GroupTableNo,
		@AccSerial,
		@OrderNo,
		@OrderType,
		@ComName,
		@CasherCode,
		@BonNo,
		@WaiterCode,
		@TableSerial,
		@CaherCashTrySerial
	)

SELECT  SCOPE_IDENTITY() OrderSerial




GO
-- attache item to order
CREATE PROCEDURE StkTr04Insert(@HeadSerial int , @ItemSerial int  , @WithMod bit, @IsMod  bit, @Qnt real = 1)
AS
BEGIN

	SET NOCOUNT ON;
	DECLARE @Price REAL
	DECLARE @MenuSerial INT
	DECLARE @ItemPosi int 
	DECLARE @OrderType int


	SELECT @OrderType = OrderType FROM StkTr03 WHERE Serial = @HeadSerial
		SELECT @MenuSerial = Serial from Menus 
		where 
		IsTakeWay = case when @Ordertype = 0 then  1 else null end 
		or  
		IsTS = case when @OrderType  = 2 then   1 else null end 
		or 
		IsDelivery = Case when @OrderType = 1 then  1 else null end 

	SELECT @ItemPosi = ISNULL(ItemPosi ,0 ) + 1 from StkTr04 where Serial = @HeadSerial 
	SELECT @Price = ItemPrice FROM ItemMenuMap WHERE ItemSerial = @ItemSerial AND MnuSerial = @MenuSerial
	INSERT INTO StkTr04 (
		HeadSerial ,
		ItemSerial,
		WithMod ,
		Price,
		Qnt,
		IsMod,
		ItemPosi

	) VALUES (
		@HeadSerial ,
		@ItemSerial,
		@WithMod,
		@price,
		@Qnt,
		@IsMod,
		@ItemPosi
	)


	
END



GO
-- remove item from order
CREATE PROCEDURE StkTr04Delete(@Serial int)
AS
BEGIN
	DELETE FROM StkTr04 WHERE StkTr04.Serial = @Serial
	SELECT 1 DeltedSucessfully
	
END





GO
-- insert into comuse table to activate the device in first time to use it
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO

CREATE PROCEDURE  [dbo].[ComUseInsert] (@Imei VARCHAR(50) , @ComName VARCHAR(100))
 AS
 
	DECLARE @Capital Char
	DECLARE @Store Char
	SELECT @Capital = CHAR(ASCII(MAX(Capital)) + 1)  FROM ComUse 
	SELECT @Store = StoreCode FROM POSOptions
	INSERT INTO ComUse (ComName , Imei , Capital , StoreCode ) VALUES (@ComName , @Imei , @Capital , @Store)



GO
-- get device from comuse table to check if device is authorized or not


SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO

CREATE PROCEDURE  [dbo].[ComUseGetDevice] (@Imei VARCHAR(50))
 AS
		SELECT  ComName , Capital , ISNULL(CashTry.Serial  , 0) CashtraySerial FROM ComUse LEFT JOIN CashTry ON ComName = ComputerName AND CloseDate IS NULL AND Paused = 0  WHERE Imei = @Imei 








GO
--list product modifers
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE PROCEDURE [dbo].[StkMs01GetModifiersBySerial](@Serial int)
AS
BEGIN
	-- SET NOCOUNT ON added to prevent extra result sets from
	-- interfering with SELECT statements.
	SET NOCOUNT ON;

	SELECT i.Serial , i.ItemCode , i.ItemName , m.ScreenNo   FROM  Modifires m JOIN StkMs01 i ON i.Serial = m.ModSerial  WHERE m.ItemSerial= @Serial

END






GO
-- check if table is already paused & pause it if its not paused
-- it will return true if the table was not paused and false if it was
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE PROCEDURE [dbo].[TablesOpenOrder](@Serial int , @EmpCode int , @Imei VARCHAR(50))
AS
BEGIN

	DECLARE @Paused BIT
	DECLARE @State VARCHAR(10)
	DECLARE @CurrentEmpCode VARCHAR(10)
	DECLARE @ComputerName VARCHAR(50)
	SELECT @Paused =  dbo.Tables.pause , @State = dbo.Tables.State From dbo.Tables WHERE Serial = @Serial


	IF @paused = 1
		BEGIN
			SELECT (0) IsOrderOpened , 'paused' msg
			RETURN
		END

	IF @State = 'Working'
		BEGIN
			SELECT @CurrentEmpCode =  WaiterCode From StkTr03 WHERE TableSerial = @Serial AND Finished = 0
			IF @CurrentEmpCode != @EmpCode
				BEGIN
					SELECT (0) IsOrerOpened , 'unauthorized' msg
					RETURN
				END
		END
		

	SELECT @ComputerName = ComName FROM ComUse WHERE Imei = @Imei
	UPDATE dbo.Tables SET pause = 1 , ComputerName = @ComputerName  WHERE Serial = @Serial 
	SELECT (1) IsOrerOpened , 'opened' msg
END







GO
-- check if table is already paused & pause it if its not paused
-- it will return true if the table was not paused and false if it was
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE PROCEDURE [dbo].[TablesUnPause](@Serial int , @Imei VARCHAR(50))
AS
BEGIN
	DECLARE @ComputerName VARCHAR(50)
	SELECT @ComputerName = ComName FROM ComUse WHERE Imei = @Imei
	UPDATE dbo.Tables SET pause = 0  WHERE Tables.Serial = @Serial  OR ComputerName = @ComputerName


	SELECT 1 IsTableClosed
END





GO
--reset the orders tables
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE PROCEDURE [dbo].[Reset]
AS
BEGIN

	TRUNCATE TABLE StkTr03;
	TRUNCATE TABLE StkTr04;
	UPDATE Tables Set pause = 0 , State = 'Free'
END






GO
--insert modifers into stktr04 
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE PROCEDURE [dbo].[Stktr04InsertModifiers] 
(@ItemsSerials nvarchar(100) ,@HeadSerial int,@OrderItemSerial  int) 
	as
	BEGIN  
	SET NOCOUNT ON;
	DECLARE @OrderType int
	declare @ItemPosi int 
	declare  @Qnt real
	declare @Price money 
	declare @MenuSerial int 
	declare  @ModifierSerial int 
--	declare @ID nvarchar(20)
	
	declare I_Serial cursor
	for
	SELECT Split.a.value('.', 'NVARCHAR(MAX)') DATA
    FROM
	  (
		 SELECT CAST('<X>'+REPLACE(@ItemsSerials, ',', '</X><X>')+'</X>' AS XML) AS String
       ) AS A
		CROSS APPLY String.nodes('/X') AS Split(a)
		
		open I_Serial
		Fetch Next From I_Serial into
		 @ModifierSerial 
	while @@FETCH_STATUS = 0
	begin 
	SELECT @OrderType = OrderType FROM StkTr03 WHERE Serial = @HeadSerial
		SELECT @MenuSerial = Serial from Menus 
		where 
		IsTakeWay = case when @Ordertype = 0 then  1 else null end 
		or  
		IsTS = case when @OrderType  = 2 then   1 else null end 
		or 
		IsDelivery = Case when @OrderType = 1 then  1 else null end 

	SELECT @ItemPosi = ISNULL(ItemPosi ,0 ) + 1 from StkTr04 where HeadSerial  = @HeadSerial 
	--SELECT @Price = ItemPrice FROM ItemMenuMap WHERE ItemSerial = @ModifierSerial AND MnuSerial = @MenuSerial
	INSERT INTO StkTr04 (
		HeadSerial ,
		ItemSerial,
		WithMod ,
		Price,
		Qnt,
		IsMod,
		ItemPosi,
		MainModSerial

	) VALUES (
		@HeadSerial ,
		@ModifierSerial,
		0,
		0,
		1,
		1,
		@ItemPosi,
		@OrderItemSerial
	)
        

		Fetch Next From I_Serial into
		 @ModifierSerial 
	end 
		Close I_Serial
		DEALLOCATE   I_Serial

		SELECT 1 Inserted
	end 








GO
-- list order and items by item serial
-- i refers to item , oi refers to orderItem
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE PROCEDURE [dbo].StkTr03ListItemsBySerial (@Serial int)
AS
BEGIN
	SELECT oi.Serial tr04Serial , Qnt , IIF(oi.IsMod = 1 , 0,Price) ItemPrice , ItemSerial , i.ItemName ,oi.IsMod,ISNULL(oi.MainModSerial , 0) MainModSerial FROM StkTr04 oi JOIN StkMs01 i ON oi.ItemSerial = i.Serial  WHERE oi.HeadSerial = @Serial
END





GO
-- update table serial on stktr03
-- close old table and open then new
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE PROCEDURE [dbo].StkTr03ChangeTable (@NewTableSerial int , @OldTableSerial int)
AS
BEGIN
	Update StkTr03 SET TableSerial = @NewTableSerial  WHERE TableSerial = @OldTableSerial AND ISNULL(TotalCash , 0) = 0
	UPDATE dbo.Tables SET pause = 0 , dbo.Tables.State = 'Free' WHERE Serial = @OldTableSerial
	UPDATE dbo.Tables SET pause = 1 , dbo.Tables.State = 'Working' WHERE Serial = @NewTableSerial
	SELECT 1 AS updated
END
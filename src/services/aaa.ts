
// protected static async getById<T extends TableName>(table: T, id: string) {
//     return await pb
//       .collection(table)
//       .getOne<selectArgs<T> | null | undefined>(id);
//   }

//   protected static async getByIdExpanded<T extends TableName>(table: T, id: string, expandArr: Array<expandableArgs<T>>) {
//     return await pb
//       .collection(table)
//       .getOne<selectArgs<T> & expandedArgs<T, typeof expandArr> | null | undefined>(id, { expand: expandArr.join(",") });
//   }

//   protected static async getSome<T extends TableName>(
//     table: T,
//     params?: GetSomeParams
//   ) {
//     return await pb.collection(table).getFullList<selectArgs<T>>(params);
//   }

//   protected static async getSomeExpanded<T extends TableName>(
//     table: T,  
//     expandArr: Array<expandableArgs<T>>,
//     params?: GetSomeParams
//   ) {
//     const query = { ...params, expand: expandArr.join(",") }
//     return await pb.collection(table).getFullList<selectArgs<T> | expandedArgs<T, typeof expandArr>>(query);
//   }


// static async getPriceById(id: string) {
//     const result = await super.getByIdExpanded("asset_prices", id, ["asset_id"]);
//     if (!result) return null;
//     return AssetPriceModel.from(result);
//   }

//   static async getForAsset(assetID: string) {
//     const result = await super.getSome("asset_prices", {
//       filter: `asset_id = "${assetID}"`,
//       sort: "logged_at",
//     });
//     return result.map(AssetPriceModel.from);
//   }

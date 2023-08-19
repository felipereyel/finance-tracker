export type AssetType =
  | "fii"
  | "federal_bond"
  | "cdb"
  | "hedge_fund"
  | "stock"
  | "other";

export const AssetTypeMap: Record<AssetType, string> = {
  fii: "FII",
  federal_bond: "Federal Bond",
  cdb: "CDB",
  hedge_fund: "Hedge Fund",
  stock: "Stock",
  other: "Other",
};

export const assetTypeOptions: AssetType[] = [
  "fii",
  "federal_bond",
  "cdb",
  "hedge_fund",
  "stock",
  "other",
];

type Metadata = {
  assets: {
    insert: {
      id?: string;
      name: string;
      type: AssetType;
      initial_price: number;
      buy_date: string;
      sell_date?: string;
      comment?: string;
    };
    update: {
      name?: string;
      comment?: string;
      sell_date?: string | null;
    };
    select: {
      id: string;
      created: string;
      updated: string;
      name: string;
      type: AssetType;
      initial_price: number;
      buy_date: string;
      sell_date?: string | null;
      comment?: string;
    };
    expandable: {};
  };
  assets_agg: {
    insert: {};
    update: {};
    select: {
      id: string;
      name: string;
      type: AssetType;
      initial_price: number;
      buy_date: string;
      sell_date?: string | null;
      comment?: string;
      latest_price: number;
      latest_date: string;
    };
    expandable: {};
  };
  asset_prices: {
    insert: {
      id?: string;
      asset_id: string;
      value: number;
      logged_at: string;
      gain: number;
      comment?: string;
    };
    update: {
      value?: number;
      logged_at?: string;
      gain?: number;
      comment?: string;
    };
    select: {
      id: string;
      created: string;
      updated: string;
      asset_id: string;
      value: number;
      logged_at: string;
      gain: number;
      comment?: string;
    };
    expandable: {
      "asset_id": "assets";
    }
  };
};

export type TableName = keyof Metadata;
type Table<T extends TableName> = Metadata[T];

// Select
type Select<T extends TableName> = Table<T>["select"];
export type selectArgs<T extends TableName> = Select<T>;

// Update
type Update<T extends TableName> = Table<T>["update"];
export type updateArgs<T extends TableName> = Partial<Update<T>>;

// Insert
type Insert<T extends TableName> = Table<T>["insert"];
export type insertArgs<T extends TableName> = Insert<T>;

// Expandable
type Expandable<T extends TableName> = Table<T>["expandable"];
export type expandableArgs<T extends TableName> = keyof Expandable<T>;
export type expandedArgs<T extends TableName, A extends Array<expandableArgs<T>>> = {
  expand: {
    [K in A[number]]: selectArgs<Metadata[T]['expandable'][K]>;
  }
}

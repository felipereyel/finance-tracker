type AssetTypeEnum = "fii" | "federal_bond" | "cdb" | "hedge_fund" | "other";

type Metadata = {
  assets: {
    insert: {
      id?: string;
      name: string;
      type: AssetTypeEnum;
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
      type: AssetTypeEnum;
      initial_price: number;
      buy_date: string;
      sell_date?: string | null;
      comment?: string;
    };
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

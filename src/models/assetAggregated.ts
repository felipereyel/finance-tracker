import { BaseModel } from "./base";
import type { selectArgs } from "./types";

export type AssetAggregatedDTO = selectArgs<"assets_agg">;

export class AssetAggregatedModel extends BaseModel<"assets_agg"> {
  private constructor(dto: AssetAggregatedDTO) {
    super("assets_agg", dto.id, dto);
  }

  static from(dto: AssetAggregatedDTO): AssetAggregatedModel {
    return new AssetAggregatedModel(dto);
  }

  static async getAssetById(id: string) {
    const result = await super.getById("assets_agg", id);
    if (!result) return null;
    return new AssetAggregatedModel(result);
  }

  static async getLiveAssets() {
    const result = await super.getSome("assets_agg", {
      filter: "sell_date = NULL",
      sort: "buy_date",
    });
    return result.map(AssetAggregatedModel.from);
  }

  get name() {
    return this.dto.name;
  }

  get type() {
    return this.dto.type;
  }

  get initialPrice() {
    return this.dto.initial_price;
  }

  get buyDate() {
    return this.dto.buy_date;
  }

  get sellDate() {
    return this.dto.sell_date;
  }

  get comment() {
    return this.dto.comment ?? "";
  }

  get latestPrice() {
    return this.dto.latest_price;
  }

  get latestDate() {
    return this.dto.latest_date;
  }
}

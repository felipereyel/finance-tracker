import { BaseModel } from "./base";
import { AssetModel } from "./assets";
import type { insertArgs, selectArgs } from "./types";

export type AssetPriceDTO = selectArgs<"asset_prices">;
export type AssetPriceCreateDTO = insertArgs<"asset_prices">;

export class AssetPriceModel extends BaseModel<"asset_prices"> {
  asset?: AssetModel

  private constructor(dto: AssetPriceDTO) {
    super("asset_prices", dto.id, dto);
  }

  static from(dto: AssetPriceDTO): AssetPriceModel {
    return new AssetPriceModel(dto);
  }

  static async create(object: AssetPriceCreateDTO): Promise<AssetPriceModel> {
    const dto = await super.insert("asset_prices", object);
    return AssetPriceModel.from(dto);
  }

  get createdAt() {
    return this.dto.created;
  }

  get updatedAt() {
    return this.dto.updated;
  }

  get assetID() {
    return this.dto.asset_id;
  }

  get value() {
    return this.dto.value;
  }

  get loggedAt() {
    return this.dto.logged_at;
  }

  get gain() {
    return this.dto.gain;
  }

  get comment() {
    return this.dto.comment || "";
  }

  set value(value) {
    this.update({ value });
  }

  set gain(gain) {
    this.update({ gain });
  }

  set loggedAt(logged_at) {
    this.update({ logged_at });
  }

  set comment(comment) {
    this.update({ comment });
  }
}

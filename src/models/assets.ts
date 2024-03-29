import { BaseModel } from "./base";
import type { insertArgs, selectArgs } from "./types";

export type AssetDTO = selectArgs<"assets">;
export type AssetCreateDTO = insertArgs<"assets">;

export class AssetModel extends BaseModel<"assets"> {
  private constructor(dto: AssetDTO) {
    super("assets", dto.id, dto);
  }

  static from(dto: AssetDTO): AssetModel {
    return new AssetModel(dto);
  }

  static async create(object: AssetCreateDTO): Promise<AssetModel> {
    const dto = await super.insert("assets", object);
    return AssetModel.from(dto);
  }

  get createdAt() {
    return this.dto.created;
  }

  get updatedAt() {
    return this.dto.updated;
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

  get walletId() {
    return this.dto.wallet;
  }

  set name(name) {
    this.update({ name });
  }

  set comment(comment) {
    this.update({ comment });
  }

  set sellDate(sell_date) {
    this.update({ sell_date });
  }
}

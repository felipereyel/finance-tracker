import { BaseModel } from "./base";
import type { insertArgs, selectArgs } from "./types";

export type WalletDTO = selectArgs<"wallets">;
export type WalletCreateDTO = insertArgs<"wallets">;

export class WalletModel extends BaseModel<"wallets"> {
  private constructor(dto: WalletDTO) {
    super("wallets", dto.id, dto);
  }

  static from(dto: WalletDTO): WalletModel {
    return new WalletModel(dto);
  }

  static async create(object: WalletCreateDTO): Promise<WalletModel> {
    const dto = await super.insert("wallets", object);
    return WalletModel.from(dto);
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

  set name(name) {
    this.update({ name });
  }
}

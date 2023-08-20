import type { insertArgs, selectArgs, TableName, updateArgs } from "./types";
import pbw from "../services/pocketbase";

export class BaseModel<T extends TableName> {
  protected constructor(
    readonly table: T,
    readonly id: string,
    protected dto: selectArgs<T>
  ) {}

  protected static async insert<
    T extends TableName,
    I extends insertArgs<T>,
  >(table: T, object: I) {
    return await pbw.createRecord(table, object);
  }

  protected async update<U extends updateArgs<T>>(config: U): Promise<any> {
    this.dto = { ...this.dto, ...config }; // revert if fail to update
    return await pbw.updateRecord(this.table, this.id, config);
  }

  async delete(): Promise<void> {
    await pbw.deleteRecord(this.table, this.id);
  }
}

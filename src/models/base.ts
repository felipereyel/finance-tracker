import type { insertArgs, selectArgs, TableName, updateArgs } from "./types";
import { pb } from "../services/pocketbase";

export class BaseModel<T extends TableName> {
  protected constructor(
    readonly table: T,
    readonly id: string,
    protected dto: selectArgs<T>
  ) {}

  protected static async insert<
    T extends TableName,
    I extends insertArgs<T>,
    S extends selectArgs<T>
  >(table: T, object: I): Promise<S> {
    return await pb.collection(table).create<S>(object);
  }

  protected async update<U extends updateArgs<T>>(config: U): Promise<any> {
    this.dto = { ...this.dto, ...config }; // revert if fail to update
    return await pb.collection(this.table).update(this.id, config);
  }

  async delete(): Promise<void> {
    await pb.collection(this.table).delete(this.id);
  }
}

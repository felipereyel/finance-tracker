import type { expandableArgs, expandedArgs, insertArgs, selectArgs, TableName, updateArgs } from "./types";
import { pb } from "../services/pocketbase";

type GetSomeParams = {
  sort?: string;
  filter?: string;
};

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

  protected static async getById<T extends TableName>(table: T, id: string) {
    return await pb
      .collection(table)
      .getOne<selectArgs<T> | null | undefined>(id);
  }

  protected static async getByIdExpanded<T extends TableName>(table: T, id: string, expandArr: Array<expandableArgs<T>>) {
    return await pb
      .collection(table)
      .getOne<selectArgs<T> & expandedArgs<T, typeof expandArr> | null | undefined>(id, { expand: expandArr.join(",") });
  }

  protected static async getSome<T extends TableName>(
    table: T,
    params?: GetSomeParams
  ) {
    return await pb.collection(table).getFullList<selectArgs<T>>(params);
  }

  protected static async getSomeExpanded<T extends TableName>(
    table: T,  
    expandArr: Array<expandableArgs<T>>,
    params?: GetSomeParams
  ) {
    const query = { ...params, expand: expandArr.join(",") }
    return await pb.collection(table).getFullList<selectArgs<T> | expandedArgs<T, typeof expandArr>>(query);
  }
}

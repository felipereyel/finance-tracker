import PocketBase from "pocketbase";
import { TableName, expandableArgs, expandedArgs, insertArgs, selectArgs, updateArgs } from "../models";

const POCKETBASE_URL = import.meta.env.VITE_POCKETBASE_URL || "/";

export const pb = new PocketBase(POCKETBASE_URL);

type QueryParams = {
  filter?: string;
  sort?: string;
}

class PBWrapper {
  pb: PocketBase;

  constructor() {
    this.pb = new PocketBase(POCKETBASE_URL);
  }

  async init() {
    return this.pb.health.check();
  }

  // Mutations

  async createRecord<
    T extends TableName,
    I extends insertArgs<T>,
    S extends selectArgs<T>
  >(table: T, object: I): Promise<S> {
    return this.pb.collection(table).create<S>(object);
  }

  async updateRecord<
    T extends TableName,
    U extends updateArgs<T>
  >(table: T, id: string, data: U): Promise<any> {
    return await this.pb.collection(table).update(id, data);
  }

  async deleteRecord(table: TableName, id: string): Promise<void> {
    await this.pb.collection(table).delete(id);
  }

  // Queries

  async getOneRecord<
    T extends TableName,
    S extends selectArgs<T>
  >(table: T, id: string): Promise<S | null | undefined> {
    return this.pb.collection(table).getOne<S | null | undefined>(id);
  }

  async getOneExpandedRecord<
    T extends TableName,
    S extends selectArgs<T> & expandedArgs<T, E>,
    E extends Array<expandableArgs<T>>
  >(table: T, id: string, expand: E): Promise<S | null | undefined> {
    return this.pb.collection(table).getOne<S | null | undefined>(id, { expand: expand.join(",") });
  }

  async getFullRecordList<
    T extends TableName,
    S extends selectArgs<T>
  >(table: T, query?: QueryParams): Promise<S[]> {
    return this.pb.collection(table).getFullList<S>(query);
  }
}

export default new PBWrapper();
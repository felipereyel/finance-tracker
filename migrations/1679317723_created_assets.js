migrate(
  (db) => {
    const collection = new Collection({
      id: "1rm3ii9vtitvzs4",
      created: "2023-03-20 13:08:43.366Z",
      updated: "2023-03-20 13:08:43.366Z",
      name: "assets",
      type: "base",
      system: false,
      schema: [
        {
          system: false,
          id: "jsjjklyf",
          name: "name",
          type: "text",
          required: true,
          unique: false,
          options: {
            min: null,
            max: null,
            pattern: "",
          },
        },
        {
          system: false,
          id: "wjzovcrs",
          name: "type",
          type: "text",
          required: true,
          unique: false,
          options: {
            min: null,
            max: null,
            pattern: "",
          },
        },
        {
          system: false,
          id: "moihprfl",
          name: "initial_price",
          type: "number",
          required: true,
          unique: false,
          options: {
            min: null,
            max: null,
          },
        },
        {
          system: false,
          id: "pa1wrgs6",
          name: "buy_date",
          type: "date",
          required: true,
          unique: false,
          options: {
            min: "",
            max: "",
          },
        },
        {
          system: false,
          id: "rpij3cv3",
          name: "sell_date",
          type: "date",
          required: false,
          unique: false,
          options: {
            min: "",
            max: "",
          },
        },
        {
          system: false,
          id: "2soliben",
          name: "comment",
          type: "text",
          required: false,
          unique: false,
          options: {
            min: null,
            max: null,
            pattern: "",
          },
        },
      ],
      listRule: "",
      viewRule: "",
      createRule: "",
      updateRule: "",
      deleteRule: "",
      options: {},
    });

    return Dao(db).saveCollection(collection);
  },
  (db) => {
    const dao = new Dao(db);
    const collection = dao.findCollectionByNameOrId("1rm3ii9vtitvzs4");

    return dao.deleteCollection(collection);
  }
);

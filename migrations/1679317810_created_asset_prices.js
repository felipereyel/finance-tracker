migrate(
  (db) => {
    const collection = new Collection({
      id: "5g9vavz477ztgq1",
      created: "2023-03-20 13:10:10.650Z",
      updated: "2023-03-20 13:10:10.650Z",
      name: "asset_prices",
      type: "base",
      system: false,
      schema: [
        {
          system: false,
          id: "kfqku1ye",
          name: "asset_id",
          type: "relation",
          required: true,
          unique: false,
          options: {
            collectionId: "1rm3ii9vtitvzs4",
            cascadeDelete: true,
            minSelect: null,
            maxSelect: 1,
            displayFields: ["name"],
          },
        },
        {
          system: false,
          id: "qp5r1rug",
          name: "value",
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
          id: "i6pf6xnn",
          name: "logged_at",
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
          id: "zpnucvgv",
          name: "gain",
          type: "number",
          required: false,
          unique: false,
          options: {
            min: null,
            max: null,
          },
        },
        {
          system: false,
          id: "gmmbwdcl",
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
    const collection = dao.findCollectionByNameOrId("5g9vavz477ztgq1");

    return dao.deleteCollection(collection);
  }
);

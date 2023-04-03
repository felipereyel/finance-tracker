migrate(
  (db) => {
    const snapshot = [
      {
        id: "_pb_users_auth_",
        created: "2023-03-20 13:02:28.641Z",
        updated: "2023-03-20 13:02:28.643Z",
        name: "users",
        type: "auth",
        system: false,
        schema: [
          {
            system: false,
            id: "users_name",
            name: "name",
            type: "text",
            required: false,
            unique: false,
            options: {
              min: null,
              max: null,
              pattern: "",
            },
          },
          {
            system: false,
            id: "users_avatar",
            name: "avatar",
            type: "file",
            required: false,
            unique: false,
            options: {
              maxSelect: 1,
              maxSize: 5242880,
              mimeTypes: [
                "image/jpeg",
                "image/png",
                "image/svg+xml",
                "image/gif",
                "image/webp",
              ],
              thumbs: null,
            },
          },
        ],
        listRule: "id = @request.auth.id",
        viewRule: "id = @request.auth.id",
        createRule: "",
        updateRule: "id = @request.auth.id",
        deleteRule: "id = @request.auth.id",
        options: {
          allowEmailAuth: true,
          allowOAuth2Auth: true,
          allowUsernameAuth: true,
          exceptEmailDomains: null,
          manageRule: null,
          minPasswordLength: 8,
          onlyEmailDomains: null,
          requireEmail: false,
        },
      },
      {
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
      },
      {
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
      },
    ];

    const collections = snapshot.map((item) => new Collection(item));

    return Dao(db).importCollections(collections, true, null);
  },
  (db) => {
    return null;
  }
);

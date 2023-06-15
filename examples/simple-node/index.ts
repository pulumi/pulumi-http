// Copyright 2019, Pulumi Corporation.  All rights reserved.

import * as http from "@pulumi/http";

export let content = http.getHttpOutput({
  url: "https://www.pulumi.com",
});

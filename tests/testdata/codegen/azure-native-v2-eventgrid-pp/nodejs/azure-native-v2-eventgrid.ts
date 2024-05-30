import * as pulumi from "@pulumi/pulumi";
import * as azure_native from "@pulumi/azure-native";

const example = new azure_native.eventgrid.EventSubscription("example", {
    expirationTimeUtc: "example",
    scope: "example",
});

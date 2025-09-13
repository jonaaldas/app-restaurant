import {
  View,
  Text,
  ScrollView,
  Pressable,
  TouchableOpacity,
  StyleSheet,
  ActivityIndicator,
  RefreshControl,
} from "react-native";
import RestaurantCard from "@/components/RestaurantCard";
import { router } from "expo-router";
import { SafeAreaView } from "react-native-safe-area-context";
import { Ionicons } from "@expo/vector-icons";
import { useRestaurantContext } from "@/app/useContext/restaurant";
import Colors from "@/constants/Colors";
import { useState } from "react";

export default function SavedRestaurants() {
  const { savedRestaurants, isLoadingSaved, refetchSavedRestaurants } = useRestaurantContext();
  const [refreshing, setRefreshing] = useState(false);

  const onRefresh = async () => {
    setRefreshing(true);
    await refetchSavedRestaurants();
    setRefreshing(false);
  };

  return (
    <SafeAreaView style={{ flex: 1 }}>
      <View
        style={{
          flexDirection: "row",
          justifyContent: "space-between",
          alignItems: "center",
          margin: 20,
        }}
      >
        <TouchableOpacity
          style={styles.closeButton}
          onPress={() => router.back()}
        >
          <Ionicons name="close" size={24} color="white" />
        </TouchableOpacity>

        <TouchableOpacity style={styles.shareButton}>
          <Ionicons name="share-outline" size={24} color="white" />
        </TouchableOpacity>
      </View>

      <Text style={{ fontSize: 24, fontWeight: "bold", margin: 20 }}>
        Saved Restaurants
      </Text>

      {isLoadingSaved ? (
        <View style={styles.loadingContainer}>
          <ActivityIndicator size="large" color="#007AFF" />
          <Text style={styles.loadingText}>Loading saved restaurants...</Text>
        </View>
      ) : (
        <>
          <Text style={{ margin: 20, marginTop: 0 }}>
            {savedRestaurants.length === 0 
              ? "No saved restaurants yet" 
              : `Found ${savedRestaurants.length} saved restaurants`}
          </Text>

          <ScrollView 
            style={{ flex: 1 }} 
            showsVerticalScrollIndicator={false}
            refreshControl={
              <RefreshControl
                refreshing={refreshing}
                onRefresh={onRefresh}
                tintColor="#007AFF"
              />
            }
          >
            {savedRestaurants.map((restaurant, index) => {
              return (
                <Pressable
                  key={restaurant.place_id}
                  onPress={() => {
                    router.push({
                      pathname: "/restaurant/[id]",
                      params: { id: restaurant.place_id },
                    });
                  }}
                >
                  <RestaurantCard restaurant={restaurant} />
                </Pressable>
              );
            })}
          </ScrollView>
        </>
      )}
    </SafeAreaView>
  );
}

const styles = StyleSheet.create({
  closeButton: {
    backgroundColor: "rgba(0, 0, 0, 0.7)",
    borderRadius: 20,
    width: 40,
    height: 40,
    justifyContent: "center",
    alignItems: "center",
  },
  shareButton: {
    backgroundColor: "rgba(0, 0, 0, 0.7)",
    borderRadius: 20,
    width: 40,
    height: 40,
    justifyContent: "center",
    alignItems: "center",
  },
  loadingContainer: {
    flex: 1,
    justifyContent: "center",
    alignItems: "center",
    paddingBottom: 100,
  },
  loadingText: {
    marginTop: 16,
    fontSize: 16,
    color: "#666",
  },
});
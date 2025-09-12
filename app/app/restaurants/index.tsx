import { View, Text, ScrollView, Pressable, TouchableOpacity, StyleSheet } from "react-native";
import RestaurantCard from "@/components/RestaurantCard";
import { Restaurant } from "@/types/restaurants";
import { useEffect, useState } from "react";
import { router } from "expo-router";
import { SafeAreaView } from "react-native-safe-area-context";
import { Ionicons } from "@expo/vector-icons";
import { useRestaurantContext } from "@/app/useContext/restaurant";

export default function Restaurants() {
  const { restaurants } = useRestaurantContext();
  console.log("Restaurants", restaurants);

  return (
    <SafeAreaView style={{ flex: 1 }}>
      <View style={{ flexDirection: 'row', justifyContent: 'space-between', alignItems: 'center', margin: 20 }}>
      <TouchableOpacity
          style={styles.closeButton}
          onPress={() => router.back()}
        >
          <Ionicons name="close" size={24} color="white" />
        </TouchableOpacity>

        <TouchableOpacity style={styles.shareButton} >
          <Ionicons name="share-outline" size={24} color="white" />
        </TouchableOpacity> 
      </View>
      <Text style={{ fontSize: 24, fontWeight: 'bold', margin: 20 }}>Restaurants</Text>
      <Text style={{ margin: 20 }}>Found {restaurants.length} restaurants</Text>
      {restaurants.map((restaurant, index) => {
        console.log('Mapping restaurant', index, restaurant.name);
        return (
          <Pressable key={restaurant.place_id} onPress={() => {
            router.push({ pathname: "/restaurant/[id]", params: { id: restaurant.place_id } });
          }}>
            <RestaurantCard restaurant={restaurant} />
          </Pressable>
        );
      })}
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
});
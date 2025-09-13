import { Restaurant, SearchParams } from "@/types/restaurants";
import { createContext, useContext, useState, ReactNode } from "react";
import { searchRestaurants } from "@/utils/restaurants";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { useRouter } from "expo-router";

interface RestaurantContextType {
  restaurants: Restaurant[];
  searchRestaurants: (params: SearchParams) => void;
  getRestaurantById: (placeId: string) => Promise<Restaurant | null>;
}

export const RestaurantContext = createContext<
  RestaurantContextType | undefined
>(undefined);

interface RestaurantProviderProps {
  children: ReactNode;
}

export const RestaurantProvider = ({ children }: RestaurantProviderProps) => {
  const [restaurants, setRestaurants] = useState<Restaurant[]>([]);
  const router = useRouter();
  const queryClient = useQueryClient();

  const searchRestaurantsMutation = useMutation({
    mutationFn: (params: SearchParams) => searchRestaurants(params),
    onSuccess: (val: Restaurant[]) => {
      queryClient.invalidateQueries({ queryKey: ["search-restaurants"] });
      setRestaurants(val);
      router.push("/restaurants");
    },
    onError: (error) => {
      console.error("Search restaurants error:", error);
    }
  });

  const getRestaurantById = async (placeId: string): Promise<Restaurant | null> => {
    // TODO: Implement
    return null;
  };

  const value: RestaurantContextType = {
    restaurants,
    searchRestaurants: searchRestaurantsMutation.mutate,
    getRestaurantById,
  };

  return (
    <RestaurantContext.Provider value={value}>
      {children}
    </RestaurantContext.Provider>
  );
};

export const useRestaurantContext = () => {
  const context = useContext(RestaurantContext);
  if (!context) {
    throw new Error(
      "useRestaurantContext must be used within a RestaurantProvider"
    );
  }
  return context;
};

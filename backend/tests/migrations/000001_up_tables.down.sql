-- keywords table
truncate keywords.word cascade;

-- recipe tables
truncate saladrecipes.user cascade;
truncate saladrecipes.saladtype cascade;
truncate saladrecipes.saladtype cascade;
truncate saladrecipes.modstatus cascade;
truncate saladrecipes.recipe cascade;
truncate saladrecipes.ingredienttype cascade;
truncate saladrecipes.ingredient cascade;
truncate saladrecipes.comment cascade;
truncate saladrecipes.measurement cascade;
truncate saladrecipes.recipestep cascade;

-- links
truncate saladrecipes.recipeingredient cascade;
truncate saladrecipes.typesofsalads cascade;

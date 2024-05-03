package blog

import (
	"pride/database"
	"pride/models"
	"pride/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection = database.GetCollection(database.DB, "blog")

func getAll(c *gin.Context) {
	page, per_page, skip, limit := utils.GetPagination(c.Request.URL.Query())

	opts := options.FindOptions{}
	opts.SetSkip(skip)
	opts.SetLimit(limit)

	cur, err := collection.Find(c, bson.M{}, &opts)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
	}

	total, err := collection.CountDocuments(c, bson.M{})
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
	}

	var blogs []models.BlogModel
	defer cur.Close(c)
	for cur.Next(c) {
		var blog models.BlogModel
		if err = cur.Decode(&blog); err != nil {
			c.JSON(500, gin.H{"message": err.Error()})
			return
		}
		blogs = append(blogs, blog)
	}

	c.JSON(200, gin.H{
		"data":       blogs,
		"pagination": bson.M{"page": page, "perPage": per_page, "total": total},
	})
}

func getOne(c *gin.Context) {
	id := c.Param("id")
	var blog models.BlogModel

	err := collection.FindOne(c, bson.M{"id": id}).Decode(&blog)
	if err != nil {
		c.JSON(404, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": blog})

}

func createOne(c *gin.Context) {

	var blog models.BlogModel
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	result, err := collection.InsertOne(c, blog)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": result})

}

func updateOne(c *gin.Context) {

	id := c.Param("id")
	var blog models.BlogModel

	err := collection.FindOne(c, bson.M{"id": id}).Decode(&blog)
	if err != nil {
		c.JSON(404, gin.H{"message": err.Error()})
		return
	}

	var updateblog models.BlogModel
	if err := c.ShouldBindJSON(&updateblog); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	update := bson.D{{Key: "$set", Value: updateblog}}

	result, err := collection.UpdateOne(c, bson.M{"id": id}, update)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": result})

}

func deleteOne(c *gin.Context) {

	id := c.Param("id")

	result, err := collection.DeleteOne(c, bson.M{"id": id})
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": result})
}
